import argparse
from utils import parse_input, input_to_field_element

try:
    parser = argparse.ArgumentParser(
                        prog='Omron AVS ezkl operator formatted')

    parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=False)
    args = parser.parse_args()

    input = (parse_input(args.input, scale=0))
    input = [input_to_field_element(float(i)) for i in input]
    output_string = ""
    for i in input:
        output_string += str(int(i)) + " "
    print(output_string)

except Exception as error:
    print(error)
