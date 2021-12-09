<script lang="ts">
  import { fade } from 'svelte/transition'
  import { goto } from '$app/navigation'
  
  let server = "http://localhost:5000"
  let username = ''
  let password = ''
  let hasMessage = ''
  let message = ''

  async function signup() {
    
    if (username.length == 0 ) {
      writeMessage('username cant be empty', 'failed')
      return
    } else if (password.length < 6) {
      writeMessage('password can\'t be shorter than 6 characters', 'failed')
      return
    }
  
    let res = await fetch(server + "/auth/signup", {
      method: 'POST',
      headers: { 
        'Content-Type': 'application/json', 
      },
      body: JSON.stringify({ username, password })
    })

    if (res.ok) {
      writeMessage(await res.text(), 'success')
      setTimeout(() => {
        goto("/login")
      }, 5000)
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

<svelte:head>
  <title>Signup</title>
</svelte:head>

<div class="section">
  <form class="box is-centered">
    <h1 class="title">Signup</h1>
    <label for="username">
      <p class="subtitle">Username</p>
      <input class="input is-primary" type="text" bind:value={username}>
    </label>
    <label for="password">
      <p class="subtitle">Password</p>
      <input class="input is-primary is-half" type="password" bind:value={password}>
    </label>
    <button class="button is-primary" on:click|preventDefault={signup}>Submit</button>
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
