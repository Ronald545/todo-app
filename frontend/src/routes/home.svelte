<script lang="ts">
  import { onMount } from 'svelte'
  import { fade } from 'svelte/transition'
  import { jwt } from '../stores/jwt'
  let server = 'http://localhost:5000'
  $: data = []
  let name = ''
  let description = ''

  let jwt_value = ''
  onMount(async () => {
    let unsub = jwt.subscribe(value => {
      jwt_value += value
      loadTasks()
    })
    unsub()
  })
  if (jwt_value == "<empty_string>") {
    alert("user not logged in")
  }

  async function loadTasks() {
    let res = await fetch(server + "/task", {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${jwt_value}`
      }
    })
    
    if (res.ok) {
      data = await res.json()  
    } else {
      console.log(await res.text())
    }
  }

  async function createTask() {
    if (name.length == 0 ) {
      console.log()
      return
    }
    let res = await fetch(server + "/task", {
      method: 'POST',
      body: JSON.stringify({ name, description }),
      headers: {
        'Authorization': `Bearer ${jwt_value}`,
        'Content-Type':'application/json'
      }
    })
    
    if (res.ok) {
      loadTasks()
    } else {
      console.log(await res.text())
    }
  }

  async function deleteTask(task : any) {

    let res = await fetch(server + `/task/${task.id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${jwt_value}`
      }
    })

    if (res.ok) {
      data.splice(data.indexOf(task), 1)
      data = [...data]
    } else {
      console.log(await res.text())
    }
  }
</script>

<div class="sector container">
  <h1 class="title is-1"> Todo List </h1>
  
  <form class="box">
    <label for="name">
        <p class="subtitle">Name of Task</p>
        <input type="text" class="input is-primary" bind:value={name}>
    </label>
    <label for="description">
        <p class="subtitle">Description</p>
        <input type="textarea" class="textarea is-primary" bind:value={description}>
    </label>
    <button class="button is-primary" on:click|preventDefault={createTask}>Submit</button>
  </form>

  {#each data as task (task.id)}
    <div class="content" transition:fade>
      <blockquote> {task.name} , {task.description} <button class="button is-danger" on:click={() => deleteTask(task)}>Delete</button></blockquote>
    </div>
  {/each}

  <button class="button is-primary" on:click={loadTasks}> Refresh </button>
</div>

<style>
  blockquote {
    height: 80px
  }
  button.is-danger {
    float: right;
    padding-top: 0
  }
  button.is-primary {
    margin-top: 15px;
    margin-left: 15px;
  }
</style>
