// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgPayPublisherRespect } from "./types/cointrunk/tx";
import { MsgAddArticle } from "./types/cointrunk/tx";


const types = [
  ["/bze.cointrunk.v1.MsgPayPublisherRespect", MsgPayPublisherRespect],
  ["/bze.cointrunk.v1.MsgAddArticle", MsgAddArticle],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgPayPublisherRespect: (data: MsgPayPublisherRespect): EncodeObject => ({ typeUrl: "/bze.cointrunk.v1.MsgPayPublisherRespect", value: MsgPayPublisherRespect.fromPartial( data ) }),
    msgAddArticle: (data: MsgAddArticle): EncodeObject => ({ typeUrl: "/bze.cointrunk.v1.MsgAddArticle", value: MsgAddArticle.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
