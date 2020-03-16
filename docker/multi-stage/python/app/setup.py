from setuptools import setup, find_packages

setup(
    name='app',
    version='1.0.3',
    description='app description',
    author='koketani',
    author_email='ykoketani@gmail.com',
    python_requires='>=3.6.0',
    url='https://github.com/koketani/playground-oss/docker/multi-stage/python/app',
    packages=find_packages(exclude=('tests',)),
    entry_points={
        'console_scripts': ['app=app.app:main'],
    }
)
