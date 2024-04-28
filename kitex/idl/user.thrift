namespace go user

include "base.thrift"

struct UserRegisterRequest {
    1: required string username;
    2: required string password (vt.min_size = "5", vt.max_size = "18");
}

struct UserRegisterResponse {
    1: base.Status base;
}

struct UserLoginRequest {
    1: required string username;
    2: required string password (vt.min_size = "5", vt.max_size = "18");
    3: string code;
}

struct UserLoginResponse {
    1: base.Status base;
    2: base.User data;
}

struct UserInfoRequest {
    1: required string user_id;
}

struct UserInfoResponse {
    1: base.Status base;
    2: base.User data;
}

struct UserAvatarUploadRequest {
    1: required string user_id;
    2: required binary data;
    3: required i64 filesize;
}

struct UserAvatarUploadResponse {
    1: base.Status base;
    2: base.User data;
}

struct AuthMfaQrcodeRequest { 
    1: string user_id;
}

struct Qrcode {
    1: string secret;
    2: string qrcode;
}
struct AuthMfaQrcodeResponse {
    1: base.Status base;
    2: Qrcode data;
}

struct AuthMfaBindRequest {
    1: required string user_id;
    2: required string code;
    3: required string secret;
}

struct AuthMfaBindResponse {
    1: base.Status base;
}

service UserService {
    UserRegisterResponse Register(1: UserRegisterRequest request);
    UserLoginResponse Login(1: UserLoginRequest request);
    UserInfoResponse Info(1: UserInfoRequest request);
    UserAvatarUploadResponse AvatarUpload(1: UserAvatarUploadRequest request);
    AuthMfaQrcodeResponse AuthMfaQrcode(1: AuthMfaQrcodeRequest request);
    AuthMfaBindResponse AuthMfaBind(1: AuthMfaBindRequest request);
}