FROM python:3.8.2

WORKDIR /src
COPY . .

# # every dependencies in /root/.local/lib/python3.8/site-packges
# RUN python setup.py install --user
# # only mypackage in /mypackage/usr/local/lib/python3.8/site-packges
# RUN python setup.py install --root /mypackage

# every dependencies in /mypackage/lib/python3.8/site-packages
ENV PYTHONPATH /mypackage/lib/python3.8/site-packages
RUN mkdir -p ${PYTHONPATH} \
 && python setup.py install --prefix /mypackage
