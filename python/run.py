import argparse
from model import Model
from utils import parse_input, input_to_field_element

try:
    parser = argparse.ArgumentParser(
                        prog='Sertn AVS ezkl operator engine')

    parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=False)
    args = parser.parse_args()
    input = parse_input(args.input)
    model = Model()
    model.eval()
    # return the answer
    # answer = 0 ## bad answer
    answer = float(model(input)[0])
    print(input_to_field_element(answer)) ## good answer
except:
    print(404)
