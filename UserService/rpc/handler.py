# -*- coding: utf-8 -*-

from user.ttypes import ActivityEvent


class UserActivityHandler:
    def __init__(self):
        self.log = {}

    def online(self, userid: int, ev: ActivityEvent):
        print("Get the message %s, %s" % (userid, ev))

    def offline(self, userid: int, ev: ActivityEvent):
        pass
