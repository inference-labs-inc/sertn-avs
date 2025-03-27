from mcl import G1, G2, GT, Fr

# modulus for the underlying field F_p of the elliptic curve
_FP_MODULUS = 21888242871839275222246405745257275088696311157297823662689037894645226208583
# modulus for the underlying field F_r of the elliptic curve
_FR_MODULUS = 21888242871839275222246405745257275088548364400416034343698204186575808495617

_FIELD_ORDER = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47

_G2_XA = 0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2
_G2_XB = 0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed
_G2_YA = 0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b
_G2_YB = 0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa

def __addmod(a, b, m):
  return (a + b) % m

def __mulmod(a, b, m):
  return(a * b) % m

def __expmod(a, b, m):
  result = 1
  base = a
  _b = b
  while _b > 0:
    # Check the least significant bit (LSB) of b
    if _b & 1:
      result = (result * base) % m
    # Right shift b by 1 (effectively dividing by 2, discarding the remainder)
    _b >>= 1
    # Square the base for the next iteration (efficient for repeated multiplication)
    base = (base * base) % m
  return result

def __g1_point(x: int, y: int) -> G1:
  res = G1()
  res.setStr(f"1 {x} {y}".encode("utf-8"))
  return res

def verify_sig(sig: G1, pub_key: G2, msg_bytes: bytes, ) -> bool:
  G2 = get_g2_generator()
  msg_point = map_to_curve(msg_bytes)
  
  gt1 = GT.pairing(msg_point, pub_key)
  gt2 = GT.pairing(sig, G2)

  return gt1 == gt2

def map_to_curve(_x: bytes) -> G1:
  beta = 0
  y = 0
  x = int.from_bytes(_x, "big") % _FP_MODULUS
  while True:
      (beta, y) = __find_y_from_x(x)
      # y^2 == beta
      if( beta == ((y * y) % _FP_MODULUS) ):
          return __g1_point(x, y)
      x = (x + 1) % _FP_MODULUS
  return __g1_point(0, 0)

def __find_y_from_x(x: int) -> 'tuple[int, int]':
  # beta = (x^3 + b) % p
  beta = __addmod(__mulmod(__mulmod(x, x, _FP_MODULUS), x, _FP_MODULUS), 3, _FP_MODULUS)
  # y^2 = x^3 + b
  # this acts like: y = sqrt(beta) = beta^((p+1) / 4)
  y = __expmod(beta, 0xc19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52, _FP_MODULUS)
  return (beta, y)

def check_g1_and_g2_discrete_log_equality(p1: G1, p2: G2) -> bool:
  G1 = get_g1_generator()
  G2 = get_g2_generator()
  
  gt1 = GT.pairing(p1, G2)
  gt2 = GT.pairing(G1, p2)
  return gt1 == gt2

def get_g1_generator() -> G1:
  g1 = G1()
  g1.setStr(b"1 1 2")
  return g1

def get_g2_generator() -> G2:
  g2 = G2()
  g2.setStr(f"1 {_G2_XB} {_G2_XA} {_G2_YB} {_G2_YA}".encode("utf-8"))
  return g2

def mul_by_generator_g1(a: Fr) -> G1:
  return get_g1_generator() * a

def mul_by_generator_g2(a: Fr) -> G2:
  return get_g2_generator() * a