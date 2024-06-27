import argparse
from importlib import import_module
from utils import parse_input, input_to_field_element
from model_db import models

parser = argparse.ArgumentParser(
                    prog='Sertn AVS ezkl operator engine')

parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=False)
parser.add_argument('-m','--model', nargs='+', help='model to run on', required=False)
args = parser.parse_args()

input = parse_input(args.input)
model_address = args.model[0]

if not (model_address in models.keys()):
    model_path = "model_0"
else:
    model_path = models[model_address]
    
model_module = import_module(model_path)

Model = getattr(model_module, "Model")

model = Model()
model.eval()
# return the answer
# answer = 0 ## bad answer
answer = float(model(input)[0])
print(input_to_field_element(answer)) ## good answer
# except Exception as err:
#     print(err)
