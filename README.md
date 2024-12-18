# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [friends.proto](#friends-proto)
    - [EmptyFriends](#-EmptyFriends)
    - [GetCountFriendsOut](#-GetCountFriendsOut)
    - [GetPeerFollowIn](#-GetPeerFollowIn)
    - [GetPeerFollowOut](#-GetPeerFollowOut)
    - [GetWhoFollowPeerIn](#-GetWhoFollowPeerIn)
    - [GetWhoFollowPeerOut](#-GetWhoFollowPeerOut)
    - [IsFriendExistIn](#-IsFriendExistIn)
    - [IsFriendExistOut](#-IsFriendExistOut)
    - [Peer](#-Peer)
    - [RemoveFriendsIn](#-RemoveFriendsIn)
    - [RemoveFriendsOut](#-RemoveFriendsOut)
    - [RemoveSubscribeIn](#-RemoveSubscribeIn)
    - [RemoveSubscribeOut](#-RemoveSubscribeOut)
    - [SetFriendsIn](#-SetFriendsIn)
    - [SetFriendsOut](#-SetFriendsOut)
    - [SetInvitePeerIn](#-SetInvitePeerIn)
    - [SetInvitePeerOut](#-SetInvitePeerOut)
  
    - [FriendsService](#-FriendsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="friends-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## friends.proto



<a name="-EmptyFriends"></a>

### EmptyFriends







<a name="-GetCountFriendsOut"></a>

### GetCountFriendsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [int64](#int64) |  |  |
| subscribers | [int64](#int64) |  |  |






<a name="-GetPeerFollowIn"></a>

### GetPeerFollowIn
Request for subscription


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Peer uuid |






<a name="-GetPeerFollowOut"></a>

### GetPeerFollowOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscription | [Peer](#Peer) | repeated |  |






<a name="-GetWhoFollowPeerIn"></a>

### GetWhoFollowPeerIn
Request for subscribers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Peer uuid |






<a name="-GetWhoFollowPeerOut"></a>

### GetWhoFollowPeerOut
Response subscribers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscribers | [Peer](#Peer) | repeated | Result of the operation |






<a name="-IsFriendExistIn"></a>

### IsFriendExistIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [string](#string) |  |  |






<a name="-IsFriendExistOut"></a>

### IsFriendExistOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="-Peer"></a>

### Peer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uuid | [string](#string) |  | Peer uuid |






<a name="-RemoveFriendsIn"></a>

### RemoveFriendsIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [string](#string) |  |  |






<a name="-RemoveFriendsOut"></a>

### RemoveFriendsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="-RemoveSubscribeIn"></a>

### RemoveSubscribeIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [string](#string) |  |  |






<a name="-RemoveSubscribeOut"></a>

### RemoveSubscribeOut







<a name="-SetFriendsIn"></a>

### SetFriendsIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [string](#string) |  |  |






<a name="-SetFriendsOut"></a>

### SetFriendsOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="-SetInvitePeerIn"></a>

### SetInvitePeerIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |






<a name="-SetInvitePeerOut"></a>

### SetInvitePeerOut






 

 

 


<a name="-FriendsService"></a>

### FriendsService
Service for friends

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetFriends | [.SetFriendsIn](#SetFriendsIn) | [.SetFriendsOut](#SetFriendsOut) | Add friends method |
| GetPeerFollow | [.GetPeerFollowIn](#GetPeerFollowIn) | [.GetPeerFollowOut](#GetPeerFollowOut) |  |
| GetWhoFollowPeer | [.GetWhoFollowPeerIn](#GetWhoFollowPeerIn) | [.GetWhoFollowPeerOut](#GetWhoFollowPeerOut) |  |
| RemoveSubscribe | [.RemoveSubscribeIn](#RemoveSubscribeIn) | [.RemoveSubscribeOut](#RemoveSubscribeOut) |  |
| SetInvitePeer | [.SetInvitePeerIn](#SetInvitePeerIn) | [.SetInvitePeerOut](#SetInvitePeerOut) |  |
| RemoveFriends | [.RemoveFriendsIn](#RemoveFriendsIn) | [.RemoveFriendsOut](#RemoveFriendsOut) |  |
| GetCountFriends | [.EmptyFriends](#EmptyFriends) | [.GetCountFriendsOut](#GetCountFriendsOut) |  |
| IsFriendExist | [.IsFriendExistIn](#IsFriendExistIn) | [.IsFriendExistOut](#IsFriendExistOut) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

