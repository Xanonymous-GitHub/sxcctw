<script setup lang="ts">
const name = ref('')

const router = useRouter()
const go = () => {
  if (name.value)
    router.push(`/hi/${encodeURIComponent(name.value)}`)
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

  <div py-4 />

  <input
    id="input"
    v-model="name"
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
      :disabled="!name"
      @click="go"
    >
      Short it
    </button>
  </div>
</template>
