<script lang="ts">
  import { onMount } from 'svelte'
  import { goto } from '$app/navigation'
  import { slide, fade } from 'svelte/transition'
  let server = 'http://localhost:5000'
  $: data = []
  let name = ''
  let description = ''

  let edit = false
  let editedTask

  let hasMessage = ''
  let message = ''

  onMount(async () => {
    loadTasks()
  })

  async function loadTasks() {
    let res = await fetch(server + "/task", {
      method: 'GET',
      credentials: 'include'
    })
    
    if (res.ok) {
      data = await res.json()  
    } else {
      let d = await res.text() 
      writeMessage(d, 'failed')     
    }
  }

  async function createTask() {
    if (name.length == 0 ) {
      writeMessage('name field cannot be empty', 'failed')
      return
    }
    let res = await fetch(server + "/task", {
      method: 'POST',
      body: JSON.stringify({ name, description }),
      headers: {
        'Content-Type':'application/json'
      },
      credentials: 'include'
    })
    
    if (res.ok) {
      const task = await res.json()
      data = [{ id: task.id, name, description }, ...data]
    } else {
      let d = await res.text()
      writeMessage(d, 'failed')
    }

    name = ""
    description = ""
  }

  async function editTask() {
    if (name.length === 0) {
      writeMessage('name field cannot be empty', 'failed')
      return 
    }
    let res = await fetch(server + '/task', {
      method: 'PUT',
      credentials: 'include',
      headers: {'Content-Type': 'application/json'} ,
      body:JSON.stringify({
        id: editedTask,
        newDescription: description,
        newName: name
      })
    })

    if (res.ok) {
      loadTasks()
    } else {
      let d = await res.text()
      writeMessage(d, 'failed')
    }

    name = ""
    description = ""
    edit = false
  }

  function toggleEditTask(task : any) {
    editedTask = task.id
    name = task.name
    description = task.description
    edit = true
  }

  async function deleteTask(task : any) {

    let res = await fetch(server + `/task/${task.id}`, {
      method: 'DELETE',
      credentials: 'include'
    })

    if (res.ok) {
      data.splice(data.indexOf(task), 1)
      data = [...data]
    } else {
      let d = await res.text()
      writeMessage(d, 'failed')
    }
  }

  async function logOut() {
    let res = await fetch(server + '/auth/logout', {
      method: 'POST',
      credentials: 'include'
    })

    if (res.ok) {
      goto("/")
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
  <title>Home Page</title>
</svelte:head>

<div class="sector container">
  <h1 class="title is-1"> Todo List </h1>
  
  <button class="button is-link" on:click|preventDefault={logOut}>Log Out</button>
  
  <form class="box">
    <label for="name">
        <p class="subtitle">Name of Task</p>
        <input type="text" class="input is-primary" bind:value={name}>
    </label>
    <label for="description">
        <p class="subtitle">Description</p>
        <input type="textarea" class="textarea is-primary" bind:value={description}>
    </label>

    {#if edit}
    <button class="button is-warning form-button" on:click|preventDefault={editTask} transition:slide>Edit</button>
    <button class="button is-link form-button" transition:slide on:click|preventDefault={() => {
      edit = false
      name = ""
      description = ""
    }}> X </button>
    {:else}
    <button class="button is-primary form-button" on:click|preventDefault={createTask} transition:slide >Submit</button>
    {/if}

    {#if hasMessage === 'failed'}
      <div class="notification is-danger">{message}</div>
    {/if}
  </form>

  {#each data as task (task.id)}
    <div class="content" in:slide out:fade>
      <blockquote> 
      <h2 class="subtitle">{task.name.toUpperCase()}</h2>
      <p>{task.description}</p>

      <button class="button is-danger" on:click={() => deleteTask(task)}>Delete</button> <button class="button is-warning" on:click={() => toggleEditTask(task)}>Edit</button></blockquote>
    </div>
  {/each}

  <button class="button is-primary" on:click={loadTasks}> Refresh </button>
</div>

<style>
  button.is-warning, button.is-danger {
    margin-left: 15px;
  }
  button.is-primary {
    margin-top: 15px;
    margin-left: 15px;
  }
  button.form-button {
    margin-top: 15px;
  }
  div.notification {
    margin-top : 15px;
  }
</style>
