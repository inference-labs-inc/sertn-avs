import argparse
from model import Model
from utils import parse_input

try:
    parser = argparse.ArgumentParser(
                        prog='Omron AVS ezkl operator engine')

    parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=False)
    args = parser.parse_args()
    input = parse_input(args.input)
    model = Model()
    model.eval()
    # return the answer
    print(int(model(input)[0]))
except:
    print(0)