// types.ts
export enum AuthProvider {
  LOCAL = 'local',
  GOOGLE = 'google',
  FACEBOOK = 'facebook',
  GITHUB = 'github',
}

export interface User {
  id: string;
  username: string;
  email: string;
  password: string;
  provider: AuthProvider;
  createdAt: Date;
  updatedAt: Date;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface TokenResponse {
  token: string;
  expiresAt: Date;
}

export interface AuthResponse {
  user: User;
  token: TokenResponse;
}