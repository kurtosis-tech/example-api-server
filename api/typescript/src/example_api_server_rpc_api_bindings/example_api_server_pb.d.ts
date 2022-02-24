import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class AddPersonArgs extends jspb.Message {
  getPersonId(): string;
  setPersonId(value: string): AddPersonArgs;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddPersonArgs.AsObject;
  static toObject(includeInstance: boolean, msg: AddPersonArgs): AddPersonArgs.AsObject;
  static serializeBinaryToWriter(message: AddPersonArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddPersonArgs;
  static deserializeBinaryFromReader(message: AddPersonArgs, reader: jspb.BinaryReader): AddPersonArgs;
}

export namespace AddPersonArgs {
  export type AsObject = {
    personId: string,
  }
}

export class GetPersonArgs extends jspb.Message {
  getPersonId(): string;
  setPersonId(value: string): GetPersonArgs;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPersonArgs.AsObject;
  static toObject(includeInstance: boolean, msg: GetPersonArgs): GetPersonArgs.AsObject;
  static serializeBinaryToWriter(message: GetPersonArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPersonArgs;
  static deserializeBinaryFromReader(message: GetPersonArgs, reader: jspb.BinaryReader): GetPersonArgs;
}

export namespace GetPersonArgs {
  export type AsObject = {
    personId: string,
  }
}

export class GetPersonResponse extends jspb.Message {
  getBooksRead(): number;
  setBooksRead(value: number): GetPersonResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPersonResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetPersonResponse): GetPersonResponse.AsObject;
  static serializeBinaryToWriter(message: GetPersonResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPersonResponse;
  static deserializeBinaryFromReader(message: GetPersonResponse, reader: jspb.BinaryReader): GetPersonResponse;
}

export namespace GetPersonResponse {
  export type AsObject = {
    booksRead: number,
  }
}

export class IncrementBooksReadArgs extends jspb.Message {
  getPersonId(): string;
  setPersonId(value: string): IncrementBooksReadArgs;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IncrementBooksReadArgs.AsObject;
  static toObject(includeInstance: boolean, msg: IncrementBooksReadArgs): IncrementBooksReadArgs.AsObject;
  static serializeBinaryToWriter(message: IncrementBooksReadArgs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IncrementBooksReadArgs;
  static deserializeBinaryFromReader(message: IncrementBooksReadArgs, reader: jspb.BinaryReader): IncrementBooksReadArgs;
}

export namespace IncrementBooksReadArgs {
  export type AsObject = {
    personId: string,
  }
}

