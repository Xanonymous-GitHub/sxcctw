<script setup lang="ts">
import { requestNewUrlShortenRecord } from '~/api/service'

const host = import.meta.env.DEV ? 'http://localhost:8080' : 'https://s.xcc.tw'

const originUrl = ref('')
const shortenedUrl = ref('')
const errMsg = ref('')

const showShortenedResult = (id: string) => {
  shortenedUrl.value = `${host}/s/${id}`
  navigator.clipboard.writeText(shortenedUrl.value)
}

const showErrMsg = (msg: string) => {
  errMsg.value = msg
}

const isValidateUrl = (url: string): boolean => {
  try {
    // eslint-disable-next-line no-new
    new URL(url)
  }
  catch (e) {
    return false
  }

  return true
}

const go = async() => {
  if (!isValidateUrl(originUrl.value))
    return
  const result = await requestNewUrlShortenRecord(originUrl.value)
  if (!result)
    return
  if (result.shortenedID)
    showShortenedResult(result.shortenedID)
  if (result.msg)
    showErrMsg(result.msg)
  else
    showErrMsg('Error occurred!')
}

</script>

<template>
  <div i-flat-color-icons-close-up-mode text-4xl inline-block />
  <h1>
    SXCCTW URL Shortener
  </h1>
  <p>
    <em text-sm op75>Shorten your URL easily and quickly, easy to use, unconstrained, safe and secure.</em>
  </p>

  <h2 v-if="shortenedUrl" text-blue my-6 select-text font-bold>
    <a :href="shortenedUrl" target="_blank" rel="noreferrer nofollow noopener">{{ shortenedUrl }}</a>
  </h2>

  <h2 v-if="!shortenedUrl && errMsg" text-red font-bold>
    {{ errMsg }}
  </h2>

  <div py-4 />

  <input
    v-if="!shortenedUrl"
    id="input"
    v-model="originUrl"
    placeholder="Origin URL?"
    type="text"
    autocomplete="false"
    p="x-4 y-2"
    w="250px"
    text="center"
    bg="transparent"
    border="~ rounded gray-200 dark:gray-700"
    outline="none active:none"
    @keydown.enter="go"
  >

  <div>
    <button
      class="m-3 text-sm btn"
      :disabled="shortenedUrl || !isValidateUrl(originUrl)"
      @click.stop="go"
    >
      {{ shortenedUrl ? 'Copied!' : 'Short it now' }}
    </button>
  </div>
</template>
