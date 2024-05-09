import argparse
import torch
from torch.autograd import Variable
from model import Model

try :
    parser = argparse.ArgumentParser(
                        prog='Omron AVS ezkl operator engine')

    parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=False)
    args = parser.parse_args()
    input = Variable(torch.Tensor([float(i) for i in args.input[0].split(" ")]))
    model = Model()

    model.eval()
    print(int(model(input)[0]))
except:
    print("10")