/**
 * Thrift files can namespace, package, or prefix their output in various
 * target languages.
 */

namespace go TransferingMess.transfer
/**
 * Structs are the basic complex data structures. They are comprised of fields
 * which each have an integer identifier, a type, a symbolic name, and an
 * optional default value.
 *
 * Fields can be declared "optional", which ensures they will not be included
 * in the serialized output if they aren't set.  Note that this requires some
 * manual management in some languages.
 */

struct ChatMessage{
    1: i64 timestamp,
    2: string toPub,
    3: string encMsg,
    4: string msgToVerifyClient,
    5: string Signature,
    6: string childPath,
    7: string toEndpoint,
}

struct Info{
    1: string addressEndpoint,
    2: string pubkeyClient,
}

service TransferMessage {
        void transfering(1:ChatMessage chatMessage),
        void eventAddClient(1: Info infoEndpoint),
        void eventRemoveClient(1: string pubkey)
}
