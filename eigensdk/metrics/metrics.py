from abc import ABC, abstractmethod


class Metrics(ABC):
    @abstractmethod
    def AddFeeEarnedTotal(amount: float, token: str): ...

    @abstractmethod
    def SetPerformanceScore(score: float): ...

    @abstractmethod
    def Start(): ...
