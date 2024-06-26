import ezkl
import argparse
from utils import relative_file_path, parse_input
import json
from model_db import models

parser = argparse.ArgumentParser(
                    prog='Sertn AVS ezkl operator proving engine')

parser.add_argument('-i','--input', nargs='+', help='input data to run on', required=True)
parser.add_argument('-m','--model', nargs='+', help='input model to run on', required=True)

args = parser.parse_args()

input = parse_input(args.input)
model_address = args.model[0]

if not (model_address in models.keys()):
    model_path = "model_0"
else:
    model_path = models[model_address]

model_path = "models/" + model_path

data_array = ((input).detach().numpy()).reshape([-1]).tolist()

data_path = relative_file_path("proof/inputs.json")
witness_path = relative_file_path("proof/witness.json")
compiled_model_path = relative_file_path(model_path + "/network.ezkl")
pk_path = relative_file_path(model_path + "/test.pk")
vk_path = relative_file_path(model_path + "/test.vk")
settings_path = relative_file_path(model_path + "/settings.json")
proof_path =    relative_file_path('proof/proof.json')

# Serialize data into file:
data = dict(input_data = [data_array])
json.dump(data, open(data_path, 'w'))


# Generate the Witness for the proof
ezkl.gen_witness(data_path, compiled_model_path, witness_path)


# Generate the proof
proof = ezkl.prove(
        witness_path,
        compiled_model_path,
        pk_path,
        proof_path,
        "single",
    )



onchain_input_array = []

# using a loop
# avoiding printing last comma

# Below is the instances but rescaling will be done on chain so all we need is a proof
formatted_output = "["
for i, value in enumerate(proof["instances"]):
    for j, field_element in enumerate(value):
        onchain_input_array.append(ezkl.felt_to_big_endian(field_element))
        formatted_output += '"' + str(onchain_input_array[-1]) + '"'
        if j != len(value) - 1:
            formatted_output += ", "
    if i != len(proof["instances"]) - 1:
        formatted_output += ", "
formatted_output += "]"

# This will be the values you use onchain
# copy them over to remix and see if they verify
# What happens when you change a value?
# print(formatted_output)

##### Output then Proof
print(onchain_input_array[-1][2:])
print(proof["proof"][2:])
