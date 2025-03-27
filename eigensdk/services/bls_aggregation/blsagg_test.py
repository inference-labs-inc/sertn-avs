import asyncio
import unittest
from eigensdk.crypto.bls.attestation import (
    KeyPair,
    Signature,
    new_g1_point,
    new_zero_g1_point,
)
from .blsagg import BlsAggregationService, BlsAggregationServiceResponse
from eigensdk.services.avsregistry.avsregistry_fake import (
    FakeAvsRegistryService,
    TestOperator,
)
from Crypto.Hash import keccak
import json


def hash_function(input_str: str):
    k = keccak.new(digest_bits=256)
    k.update(input_str.encode())
    # return k.hexdigest() # return hex string
    return k.digest()


async def run_delayed(delay, instance, method_name, *args):
    method = getattr(instance, method_name)
    if not asyncio.iscoroutinefunction(method):
        raise TypeError("coroutine must be an awaitable object")
    await asyncio.sleep(delay)
    await method(*args)


class TestBlsAggregationService(unittest.IsolatedAsyncioTestCase):
    time_to_expire_task = 3  # secound

    async def test_case_1(self):
        """1 quorum 1 operator 1 correct signature"""

        task_index = 1
        block_number = 1
        task_response = "sample text response"
        task_response_digest = hash_function(task_response)
        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        bls_sign = operator_1.bls_key_pair.sign_message(task_response_digest)
        fake_avs_registry_service = FakeAvsRegistryService(block_number, [operator_1])

        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service,
            hash_function,
        )

        await bls_aggregation_service.initialize_new_task(
            task_index=task_index,
            task_created_block=1,
            quorum_numbers=[1],
            quorum_threshold_percentages=[100],
            time_to_expiry=self.time_to_expire_task,
        )

        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index=task_index,
                task_response=task_response,
                bls_sign=bls_sign,
                operator_id=operator_1.operator_id,
            )
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index,
            task_response=task_response,
            task_response_digest=task_response_digest,
            non_signers_pubkeys_g1=[],
            quorum_apks_g1=[operator_1.bls_key_pair.pub_g1],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2,
            signers_agg_sig_g1=bls_sign,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)

        self.assertIsNone(err, f"No error expected.")
        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_2(self):
        """1 quorum 3 operator 3 correct signatures"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        operator_2 = TestOperator(
            operator_id=2,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("02"),
        )

        operator_3 = TestOperator(
            operator_id=3,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("03"),
        )

        block_number = 1
        task_index = 1
        quorum_numbers = [1]
        quorum_threshold_percentages = [100]
        task_response = "sample text response for tast case 2"
        task_response_digest = hash_function(task_response)
        fake_avs_registry_service = FakeAvsRegistryService(
            block_number, [operator_1, operator_2, operator_3]
        )

        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        await bls_aggregation_service.initialize_new_task(
            task_index,
            block_number,
            quorum_numbers,
            quorum_threshold_percentages,
            self.time_to_expire_task,
        )

        sign_1 = operator_1.bls_key_pair.sign_message(task_response_digest)
        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index, task_response, sign_1, operator_1.operator_id
            )
        )

        sign_2 = operator_2.bls_key_pair.sign_message(task_response_digest)
        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index, task_response, sign_2, operator_2.operator_id
            )
        )

        sign_3 = operator_3.bls_key_pair.sign_message(task_response_digest)
        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index, task_response, sign_3, operator_3.operator_id
            )
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index,
            task_response=task_response,
            task_response_digest=task_response_digest,
            non_signers_pubkeys_g1=[],
            quorum_apks_g1=[
                operator_1.bls_key_pair.pub_g1
                + operator_2.bls_key_pair.pub_g1
                + operator_3.bls_key_pair.pub_g1
            ],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2
            + operator_2.bls_key_pair.pub_g2
            + operator_3.bls_key_pair.pub_g2,
            signers_agg_sig_g1=sign_1 + sign_2 + sign_3,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)

        self.assertIsNone(err, f"No error expected.")
        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_3(self):
        """2 quorums 2 operators 2 correct signatures"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        operator_2 = TestOperator(
            operator_id=2,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("02"),
        )

        block_number = 1
        task_index = 1
        quorum_numbers = [1, 2]
        quorum_threshold_percentages = [100, 100]
        task_response = "sample text response for tast case 3"
        task_response_digest = hash_function(task_response)
        fake_avs_registry_service = FakeAvsRegistryService(
            block_number, [operator_1, operator_2]
        )

        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        await bls_aggregation_service.initialize_new_task(
            task_index,
            block_number,
            quorum_numbers,
            quorum_threshold_percentages,
            self.time_to_expire_task,
        )

        sign_1 = operator_1.bls_key_pair.sign_message(task_response_digest)
        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index, task_response, sign_1, operator_1.operator_id
            )
        )

        sign_2 = operator_2.bls_key_pair.sign_message(task_response_digest)
        asyncio.create_task(
            bls_aggregation_service.process_new_signature(
                task_index, task_response, sign_2, operator_2.operator_id
            )
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index,
            task_response=task_response,
            task_response_digest=task_response_digest,
            non_signers_pubkeys_g1=[],
            quorum_apks_g1=[
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
            ],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2
            + operator_2.bls_key_pair.pub_g2,
            signers_agg_sig_g1=sign_1 + sign_2,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)

        self.assertIsNone(err, f"No error expected.")
        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_4(self):
        """2 concurrent tasks 2 quorums 2 operators 2 correct signatures"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        operator_2 = TestOperator(
            operator_id=2,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("02"),
        )

        block_number = 1
        quorum_numbers = [1, 2]
        quorum_threshold_percentages = [100, 100]

        fake_avs_registry_service = FakeAvsRegistryService(
            block_number, [operator_1, operator_2]
        )

        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        task_index_1 = 1
        task_response_1 = "sample text response for tast case 4.1"
        task_response_digest_1 = hash_function(task_response_1)
        await bls_aggregation_service.initialize_new_task(
            task_index_1,
            block_number,
            quorum_numbers,
            quorum_threshold_percentages,
            self.time_to_expire_task,
        )

        task_index_2 = 2
        task_response_2 = "sample text response for tast case 4.2"
        task_response_digest_2 = hash_function(task_response_2)
        await bls_aggregation_service.initialize_new_task(
            task_index_2,
            block_number,
            quorum_numbers,
            quorum_threshold_percentages,
            self.time_to_expire_task,
        )

        task_1_sign_1 = operator_1.bls_key_pair.sign_message(task_response_digest_1)
        asyncio.create_task(
            run_delayed(
                1.0,
                bls_aggregation_service,
                "process_new_signature",
                task_index_1,
                task_response_1,
                task_1_sign_1,
                operator_1.operator_id,
            )
        )

        task_2_sign_1 = operator_1.bls_key_pair.sign_message(task_response_digest_2)
        asyncio.create_task(
            run_delayed(
                1.2,
                bls_aggregation_service,
                "process_new_signature",
                task_index_2,
                task_response_2,
                task_2_sign_1,
                operator_1.operator_id,
            )
        )

        task_1_sign_2 = operator_2.bls_key_pair.sign_message(task_response_digest_1)
        asyncio.create_task(
            run_delayed(
                1.3,
                bls_aggregation_service,
                "process_new_signature",
                task_index_1,
                task_response_1,
                task_1_sign_2,
                operator_2.operator_id,
            )
        )

        task_2_sign_2 = operator_2.bls_key_pair.sign_message(task_response_digest_2)
        asyncio.create_task(
            run_delayed(
                1.2,
                bls_aggregation_service,
                "process_new_signature",
                task_index_2,
                task_response_2,
                task_2_sign_2,
                operator_2.operator_id,
            )
        )

        want_aggregated_response_1 = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index_1,
            task_response=task_response_1,
            task_response_digest=task_response_digest_1,
            non_signers_pubkeys_g1=[],
            quorum_apks_g1=[
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
            ],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2
            + operator_2.bls_key_pair.pub_g2,
            signers_agg_sig_g1=task_1_sign_1 + task_1_sign_2,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        want_aggregated_response_2 = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index_2,
            task_response=task_response_2,
            task_response_digest=task_response_digest_2,
            non_signers_pubkeys_g1=[],
            quorum_apks_g1=[
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1,
            ],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2
            + operator_2.bls_key_pair.pub_g2,
            signers_agg_sig_g1=task_2_sign_1 + task_2_sign_2,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        (
            got_aggregated_response_1,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index_1)
        self.assertIsNone(err, f"No error expected.")
        (
            got_aggregated_response_2,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index_2)
        self.assertIsNone(err, f"No error expected.")

        self.assertEqual(
            want_aggregated_response_1.to_json(),
            got_aggregated_response_1.to_json(),
            "Want response and got response must be equal.",
        )

        self.assertEqual(
            want_aggregated_response_2.to_json(),
            got_aggregated_response_2.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_5(self):
        """1 quorum 1 operator 0 signatures - task expired"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        task_index = 1
        block_number = 1

        fake_avs_registry_service = FakeAvsRegistryService(block_number, [operator_1])
        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        await bls_aggregation_service.initialize_new_task(
            task_index=task_index,
            task_created_block=1,
            quorum_numbers=[1],
            quorum_threshold_percentages=[100],
            time_to_expiry=self.time_to_expire_task,
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=asyncio.TimeoutError(f"task {task_index} expired")
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)
        print(got_aggregated_response.to_json())

        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_6(self):
        """1 quorum 2 operator 1 correct signature quorumThreshold 50% - verified"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        operator_2 = TestOperator(
            operator_id=2,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("02"),
        )

        task_index = 1
        quorum_numbers = [1]
        quorum_threshold_percentages = [50]
        task_response = "sample text response for tast case 6"
        task_response_digest = hash_function(task_response)
        bls_sign = operator_1.bls_key_pair.sign_message(task_response_digest)
        block_number = 1

        fake_avs_registry_service = FakeAvsRegistryService(
            block_number, [operator_1, operator_2]
        )
        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        await bls_aggregation_service.initialize_new_task(
            task_index=task_index,
            task_created_block=block_number,
            quorum_numbers=quorum_numbers,
            quorum_threshold_percentages=quorum_threshold_percentages,
            time_to_expiry=self.time_to_expire_task,
        )

        asyncio.create_task(
            run_delayed(
                1.5,
                bls_aggregation_service,
                "process_new_signature",
                task_index,
                task_response,
                bls_sign,
                operator_1.operator_id,
            )
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=None,
            task_index=task_index,
            task_response=task_response,
            task_response_digest=task_response_digest,
            non_signers_pubkeys_g1=[operator_2.bls_key_pair.pub_g1],
            quorum_apks_g1=[
                operator_1.bls_key_pair.pub_g1 + operator_2.bls_key_pair.pub_g1
            ],
            signers_apk_g2=operator_1.bls_key_pair.pub_g2,
            signers_agg_sig_g1=bls_sign,
            non_signer_quorum_bitmap_indices=[],
            quorum_apk_indices=[],
            total_stake_indices=[],
            non_signer_stake_indices=[],
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)

        self.assertIsNone(err, f"No error expected.")
        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    async def test_case_7(self):
        """1 quorum 2 operator 1 correct signature quorumThreshold 60% - task expired"""

        operator_1 = TestOperator(
            operator_id=1,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("01"),
        )

        operator_2 = TestOperator(
            operator_id=2,
            stake_per_quorum={1: 100, 2: 200},
            bls_key_pair=KeyPair.from_string("02"),
        )

        task_index = 1
        quorum_numbers = [1]
        quorum_threshold_percentages = [60]
        task_response = "sample text response for tast case 7"
        task_response_digest = hash_function(task_response)
        bls_sign = operator_1.bls_key_pair.sign_message(task_response_digest)
        block_number = 1

        fake_avs_registry_service = FakeAvsRegistryService(
            block_number, [operator_1, operator_2]
        )
        bls_aggregation_service = BlsAggregationService(
            fake_avs_registry_service, hash_function
        )

        await bls_aggregation_service.initialize_new_task(
            task_index=task_index,
            task_created_block=block_number,
            quorum_numbers=quorum_numbers,
            quorum_threshold_percentages=quorum_threshold_percentages,
            time_to_expiry=self.time_to_expire_task,
        )

        asyncio.create_task(
            run_delayed(
                1.5,
                bls_aggregation_service,
                "process_new_signature",
                task_index,
                task_response,
                bls_sign,
                operator_1.operator_id,
            )
        )

        want_aggregated_response = BlsAggregationServiceResponse(
            err=asyncio.TimeoutError(f"task {task_index} expired")
        )

        (
            got_aggregated_response,
            err,
        ) = await bls_aggregation_service.get_aggregated_response(task_index)

        self.assertIsNone(err, f"No error expected.")
        self.assertEqual(
            want_aggregated_response.to_json(),
            got_aggregated_response.to_json(),
            "Want response and got response must be equal.",
        )

    # async def test_case_8(self):
    # 	"""2 quorums 2 operators which just stake one quorum; 2 correct signature - verified"""

    # async def test_case_9(self):
    # 	"""2 quorums 3 operators which just stake one quorum; 2 correct signature quorumThreshold 50% - verified"""

    # async def test_case_10(self):
    # 	"""2 quorums 3 operators which just stake one quorum; 2 correct signature quorumThreshold 60% - task expired"""

    # async def test_case_11(self):
    # 	"""2 quorums 1 operators which just stake one quorum; 1 signatures - task expired"""

    # async def test_case_12(self):
    # 	"""2 quorums 2 operators, 1 operator which just stake one quorum; 1 signatures - task expired"""

    # async def test_case_13(self):
    # 	"""send signature of task that isn't initialized - task not found error"""

    # async def test_case_14(self):
    # 	"""send new signedTaskDigest before listen on responseChan - context timeout cancels the request to prevent deadlock"""

    # async def test_case_15(self):
    # 	"""1 quorum 2 operator 2 signatures on 2 different msgs - task expired"""

    # async def test_case_16(self):
    # 	"""1 quorum 1 operator 1 invalid signature (TaskResponseDigest does not match TaskResponse)"""
