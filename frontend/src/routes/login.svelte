<script lang="ts">
  import { fade } from 'svelte/transition'
  import { goto } from '$app/navigation';
  let server = "http://localhost:5000"
  let username = ''
  let password = ''
  let hasMessage = ''
  let message = ''

  async function login() {
    
    if (username.length == 0 || password.length == 0) {
      writeMessage('username and password cant be empty', 'failed')
      return
    }
    
    let res = await fetch(server + "/auth/login", {
      method: 'POST',
      headers: { 
        'Content-Type': 'application/json', 
      },
      credentials: 'include',
      body: JSON.stringify({ username, password })
    })

    if (res.status == 200) {
      let data = await res.json()
      writeMessage('user sucessfully logged in', 'success')
      goto('/home')
      return
    } else {
      writeMessage(await res.text(), 'failed')
      return
    }
  }
  
  function writeMessage (text: string, stat: string) {
    hasMessage = stat
    message = text
    setTimeout(() => {
      hasMessage = ''
    }, 5000)
  }
</script>

<div class="section">
  <form class="box is-centered">
    <h1 class="title">Login</h1>
    <label for="username">
      <p class="subtitle">Username</p>
      <input class="input is-primary" type="text" bind:value={username}>
    </label>
    <label for="password">
      <p class="subtitle">Password</p>
      <input class="input is-primary is-half" type="password" bind:value={password}>
    </label>
    <button class="button is-primary" on:click|preventDefault={login}>Submit</button>
  </form>

  {#if hasMessage == "success"}
    <div transition:fade class="notification is-primary"> {message} </div>
  {:else if hasMessage == "failed"}
    <div transition:fade class="notification is-danger"> {message} </div>
  {/if}
</div>

<style>
  input {
    margin-bottom: 20px;
  }
</style>
