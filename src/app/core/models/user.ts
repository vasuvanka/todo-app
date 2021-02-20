export class User {
    email: string;
    name: string;
    role: string;
    id: string;
    when: string;
}

export class JwtPayload {
    name: string;
    role: string;
    id: string;
}

export class UserSignUp {
    email: string;
    password: string;
    name: string;
}

export class LoginSuccess {
    token: string;
}

export class GeneralResponse {
    status?: number;
    message: string;
}
