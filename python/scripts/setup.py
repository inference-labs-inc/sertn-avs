from torch.autograd import Variable
import ezkl
import torch
import os
import json
import sys
sys.path.append(os.path.join(os.path.dirname(__file__), "../"))

from utils import relative_file_path
from model import Model

model = Model()

model_path = relative_file_path('model_data/network.onnx')
compiled_model_path = relative_file_path('model_data/network.ezkl')
pk_path = relative_file_path('model_data/test.pk')
vk_path = relative_file_path('model_data/test.vk')
settings_path = relative_file_path('model_data/settings.json')

witness_path = relative_file_path('model_data/witness.json')
data_path = relative_file_path('model_data/input.json')
cal_data_path = relative_file_path('model_data/cal_data.json')

# Flips the neural net into inference mode
model.eval()

example_input = Variable(torch.Tensor([1,10,1,1,2]))

example_input = example_input.reshape(1, 5)


# Export the model
torch.onnx.export(model,                     # model being run
                  example_input,                         # model input (or a tuple for multiple inputs)
                  model_path,                # where to save the model (can be a file or file-like object)
                  export_params=True,        # store the trained parameter weights inside the model file
                  opset_version=10,          # the ONNX version to export the model to
                  do_constant_folding=True,  # whether to execute constant folding for optimization
                  input_names = ['input'],   # the model's input names
                  output_names = ['output'], # the model's output names
                  dynamic_axes={'input' : {0 : 'batch_size'},    # variable length axes
                                'output' : {0 : 'batch_size'}})

data_array = ((example_input).detach().numpy()).reshape([-1]).tolist()

# Edit settings
py_run_args = ezkl.PyRunArgs()
py_run_args.input_visibility = "public"
py_run_args.output_visibility = "public"
py_run_args.param_visibility = "fixed"

res = ezkl.gen_settings(model_path, settings_path, py_run_args = py_run_args)

# use the test set to calibrate the circuit
cal_data = dict(input_data = example_input.tolist())

# Serialize calibration data into file:
json.dump(cal_data, open(cal_data_path, 'w'))




# Optimize for resources, we cap logrows at 12 to reduce setup and proving time, at the expense of accuracy
# You may want to increase the max logrows if accuracy is a concern
res = ezkl.calibrate_settings(cal_data_path, model_path, settings_path, "resources", max_logrows = 12, scales = [2])

res = ezkl.compile_circuit(model_path, compiled_model_path, settings_path)

res = ezkl.get_srs(settings_path)

res = ezkl.setup(
        compiled_model_path,
        vk_path,
        pk_path,
    )
