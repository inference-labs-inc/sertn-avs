import unittest
from eigensdk.crypto.bls.attestation import KeyPair, PrivateKey


class TestAttestation(unittest.TestCase):
    def test_pk_from_string(self):
        """PrivateKey.from_string should work"""

        priv_key_str = (
            "0000000000000000000000000000000000000000000000000000000012345678"
        )
        key_pair = KeyPair.from_string(priv_key_str)
        self.assertEqual(key_pair.priv_key.get_str(), priv_key_str)

    def test_save_and_load(self):
        """Save and load of keystore should work"""

        priv_key_str = (
            "0000000000000000000000000000000000000000000000000000000012345678"
        )
        key_pair = KeyPair.from_string(priv_key_str)

        password = b"123"
        path_to_save = "./test-keystore-save/keystore.json"
        key_pair.save_to_file(path_to_save, password)
        kp2 = KeyPair.read_from_file(path_to_save, password)
        self.assertEqual(priv_key_str, kp2.priv_key.get_str())
