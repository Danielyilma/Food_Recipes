export interface User {
  username: string;
  email: string;
  profile_image: null;
  phone_number: string;
}

export interface LoginData {
  login: Login;
}

export interface Login {
  token: string;
}

export interface CustomJwtPayload {
  user_id: string;
}
