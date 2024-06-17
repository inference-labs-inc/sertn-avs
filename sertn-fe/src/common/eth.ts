import { Eip1193Provider, Log, ethers } from "ethers";

type EthWindow = {
  ethereum: Eip1193Provider;
};

const getEthWindow = () => {
  return window as unknown as EthWindow;
};

export const connectWallet = async (setSigner: any) => {
  const ethWindow = getEthWindow();
  const provider = new ethers.BrowserProvider(ethWindow.ethereum);
  const signer = await provider.getSigner();

  setSigner(signer);
};

export const getTaskManagerContract = () => {};

export const getZkVerifierContract = () => {};

export const connectContract = () => {};

export const getLogs = async (
  contract: ethers.Contract,
  signer: ethers.Signer | null
): Promise<(Log | ethers.EventLog)[]> => {
  const events = await contract.queryFilter(
    "*",
    0,
    await signer?.provider?.getBlockNumber()
  );
  console.log({ events });
  return events;
};

export const formatAddress = (address: string): string => {
  return (
    address.substring(0, 6) + "..." + address.substring(address.length - 6)
  );
};
