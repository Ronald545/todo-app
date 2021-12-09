<script lang="ts">
  import { onMount } from 'svelte'
  import { goto } from '$app/navigation'
  let server = "http://localhost:5000"

  onMount(async () => {
    let res = await fetch(server + "/task/verify", {
      credentials: 'include'
    })
    if (res.ok) {
      goto("/home")
      return
    }
  })
</script>

<svelte:head>
  <title>Welcome !</title>
</svelte:head>

<div class="container section">
  <div class="columns">
    <div class="column is-one-quarter">
      <img src="td-logo.png" alt="logo">
    </div>
    <div class="column">
      <h1 class="title is-1"> TODO APP </h1> 
      <h2 class="subtitle"> Written By Ronald Pang in Go and Svelte</h2>

      <a class="button is-link" href="/signup"> Signup </a>
      <a class="button is-link" href="/login"> Login </a>
    </div>
  </div>

</div>
