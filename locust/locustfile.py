import datetime
import random
import time

# from locust import SequentialTaskSet, task, User, constant, StopUser
from locust import SequentialTaskSet, task, User, constant
from locust.exception import StopUser


# print(dir(locust.exception))


class SleepTask(SequentialTaskSet):

    wait_time = constant(1)

    @task
    def make_bed(self):
        elapsed = random.randint(2, 5)
        time.sleep(elapsed)
        self.user.environment.events.request_success.fire(
            request_type='make', name='make', response_time=elapsed, response_length=0,
        )

    @task
    def wash_teeth(self):
        elapsed = random.randint(2, 5)
        time.sleep(elapsed)
        self.user.environment.events.request_success.fire(
            request_type='wash', name='wash', response_time=elapsed, response_length=0,
        )

    @task
    def fall_in_sleep(self):
        elapsed = random.randint(8, 10)
        time.sleep(elapsed)
        self.user.environment.events.request_success.fire(
            request_type='sleep', name='sleep', response_time=elapsed, response_length=0,
        )

    @task
    def done(self):
        # raise StopUser()
        # pass
        # raise StopLocust()
        time.sleep(1)
        self.interrupt()
        # raise StopUser()
        # self.interrupt()

    def on_start(self):
        print(f"{datetime.datetime.now()} start")

    def on_stop(self):
        print(f"{datetime.datetime.now()} stop")


class SleepyUser(User):
    host = "localhost"
    # no meaning...
    wait_time = constant(10)
    tasks = [SleepTask]

    # def __init__(self, *args, **kwargs):
    #     super().__init__(*args, **kwargs)
    #     self.client = Client(self.host)
    #     self.client._locust_environment = self.environment

    # @task
    # def sleep(self):
    #     self.client.sleep(random.randint(8, 10))

    # @task
    # def sleep2(self):
    #     self.client.sleep2(random.randint(8, 10))
