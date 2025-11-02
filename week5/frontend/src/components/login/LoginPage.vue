<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const router = useRouter();

async function login() {
  loading.value = true;
  error.value = "";

  try {
    const res = await fetch(`${import.meta.env.VITE_DIRECTUS_API}/auth/login`, {
      method: "POST",
      credentials: "include",
      mode: "cors",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
        mode: "session",
      }),
    });

    if (res.status == 400) {
      error.value = "Missing input";
      return;
    }

    if (res.status != 200) {
      error.value = "Wrong credentials";
      return;
    }

    router.push({ path: "/" });
  } catch {
    error.value = "Server error";
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div
    class="flex h-fit min-h-screen w-full flex-col items-center justify-center p-6"
  >
    <div
      class="flex w-full max-w-2xl flex-col gap-4 rounded-lg border border-zinc-300 p-6 shadow-xl"
    >
      <h1 class="text-center text-xl font-bold">Login</h1>
      <div class="flex w-full flex-col gap-2">
        <label class="flex w-full flex-col gap-1">
          Email
          <input
            type="email"
            name="email"
            class="rounded-lg border border-zinc-300 px-4 py-2 duration-200 outline-none placeholder:text-black/50 hover:border-zinc-500 active:ring-1 active:ring-blue-500"
            placeholder="johndoe@example.com"
            v-model="email"
          />
        </label>

        <label class="flex w-full flex-col gap-1">
          Password
          <input
            type="password"
            name="password"
            class="rounded-lg border border-zinc-300 px-4 py-2 duration-200 outline-none placeholder:text-black/50 hover:border-zinc-500 active:ring-1 active:ring-blue-500"
            v-model="password"
          />
        </label>
      </div>

      <p class="w-full text-right text-red-600" v-if="error">{{ error }}</p>

      <button
        class="w-fit items-center self-end rounded-lg border border-black bg-black px-4 py-2 text-white duration-200 hover:bg-transparent hover:text-black disabled:cursor-progress disabled:opacity-50"
        @click="login"
        :disabled="loading"
      >
        {{ loading ? "Logging in..." : "Login" }}
      </button>
    </div>
  </div>
</template>
