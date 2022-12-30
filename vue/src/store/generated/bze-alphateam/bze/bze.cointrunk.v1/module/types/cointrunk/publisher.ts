/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "bze.cointrunk.v1";

export interface Publisher {
  name: string;
  address: string;
  active: boolean;
  articles_count: number;
  created_at: number;
  respect: number;
}

const basePublisher: object = {
  name: "",
  address: "",
  active: false,
  articles_count: 0,
  created_at: 0,
  respect: 0,
};

export const Publisher = {
  encode(message: Publisher, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.active === true) {
      writer.uint32(24).bool(message.active);
    }
    if (message.articles_count !== 0) {
      writer.uint32(32).uint32(message.articles_count);
    }
    if (message.created_at !== 0) {
      writer.uint32(40).int64(message.created_at);
    }
    if (message.respect !== 0) {
      writer.uint32(48).int64(message.respect);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Publisher {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePublisher } as Publisher;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.active = reader.bool();
          break;
        case 4:
          message.articles_count = reader.uint32();
          break;
        case 5:
          message.created_at = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.respect = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Publisher {
    const message = { ...basePublisher } as Publisher;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = Boolean(object.active);
    } else {
      message.active = false;
    }
    if (object.articles_count !== undefined && object.articles_count !== null) {
      message.articles_count = Number(object.articles_count);
    } else {
      message.articles_count = 0;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = Number(object.created_at);
    } else {
      message.created_at = 0;
    }
    if (object.respect !== undefined && object.respect !== null) {
      message.respect = Number(object.respect);
    } else {
      message.respect = 0;
    }
    return message;
  },

  toJSON(message: Publisher): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.address !== undefined && (obj.address = message.address);
    message.active !== undefined && (obj.active = message.active);
    message.articles_count !== undefined &&
      (obj.articles_count = message.articles_count);
    message.created_at !== undefined && (obj.created_at = message.created_at);
    message.respect !== undefined && (obj.respect = message.respect);
    return obj;
  },

  fromPartial(object: DeepPartial<Publisher>): Publisher {
    const message = { ...basePublisher } as Publisher;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = object.active;
    } else {
      message.active = false;
    }
    if (object.articles_count !== undefined && object.articles_count !== null) {
      message.articles_count = object.articles_count;
    } else {
      message.articles_count = 0;
    }
    if (object.created_at !== undefined && object.created_at !== null) {
      message.created_at = object.created_at;
    } else {
      message.created_at = 0;
    }
    if (object.respect !== undefined && object.respect !== null) {
      message.respect = object.respect;
    } else {
      message.respect = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
