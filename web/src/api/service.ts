import type { AxiosError } from 'axios'
import axios from 'axios'

const urlShortenServiceApiPath = '/api/url'
axios.defaults.baseURL = import.meta.env.DEV ? 'http://localhost:8080' : undefined

export interface ShortenResponse {
  shortenedID?: string
  msg?: string
}

export const requestNewUrlShortenRecord = async(originUrl: string): Promise<ShortenResponse | null> => {
  try {
    const { data } = await axios.post<ShortenResponse>(urlShortenServiceApiPath, {
      originUrl,
    })

    return data
  }
  catch (e) {
    const err = e as AxiosError<ShortenResponse>
    return err.response?.data ?? null
  }
}
