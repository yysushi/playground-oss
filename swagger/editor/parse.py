import yaml

with open("openapi.yaml") as f:
    openapi = yaml.load(f.read())
print(openapi)
