import { API_PREFIX } from '@/constants/api'

export function buildApiUrl({ path }: { path: string }): string {
  return process.env.BACKEND_URL + API_PREFIX + path
}

export function buildPublicApiUrl({ path }: { path: string }): string {
  return process.env.NEXT_PUBLIC_BACKEND_URL + API_PREFIX + path
}
