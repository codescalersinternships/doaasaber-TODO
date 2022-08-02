<script>
    import { onMount } from "svelte/internal";
    
    let todos = []
    const baseURL = "http://localhost:8080/todo"
    
    const getTodos = async () => {
      const response = await fetch(baseURL)
      const data = await response.json()
      todos = await data
    }
    
    onMount(() => {
      getTodos()
    })
    
    
    let Task
    let ID
    const createTodo = async event => {
      event.preventDefault() 
      await fetch(baseURL, {
        method: "post",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: ID,
          task: Task,
        }),
      })
    
      getTodos()
      ID=""
      Task = ""
    }
    
    const deleteTodo = async (todo) => {
      event.preventDefault()
      await fetch(baseURL+"/"+todo.id,{
        method: "delete",
      })
      getTodos()
    }
    
    const markTask = async (todo) => {
      event.preventDefault()
      await fetch(baseURL + "/"+todo.id,{
        method: "patch",
      })
      getTodos()
    }
    
    
    </script>
    
    <main>
      <h2>TODO List</h2>
    <div class="center">
      <form on:submit="{createTodo}">
        <div class="txt_filed">
        <label>ID</label> 

        <input type="text" bind:value="{ID}" id="inputID" name="ID">
        <label>Task</label> 

        <input type="text" bind:value="{Task}" id="inputTask" name="Task">
        </div>
        <input type="submit" value="Submit" class="button"/>
      </form>
    
      <hr  style="color: #BC3CBC" >
      <hr  style="color: #BC3CBC" >
    
      {#each todos as todo}
      <div>
        <input type="checkbox" on:change={(e) => markTask(todo)}>
        <h3>{todo.id} {todo.task}</h3>
        <button name="btn" class="btn" on:click={(e) => deleteTodo(todo)}><i calss="fa fa-trash">❌</i></button>
      </div>
      {/each}
    </div>
    </main>
    
      
    
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <style>

     .center{
      position: absolute;
      top: 40%;
      left: 50%;
      transform: translate(-50%,-50%);
      width: 400px;
      background: white;
      border-radius: 10px ;

    }
   
    
      input[type=text]{
        margin: 8px 0;
        border-radius: 20px;
    
      }
    
      #inputID{
        width: 70%;
        box-sizing: border-box;
        padding: 12px 20px;
        border: 1px solid rgb(20, 64, 209);
      }
    
      #inputTask{
        width: 70%;
        box-sizing: border-box;
        padding: 12px 20px;
        border: 1px solid rgb(20, 64, 209);
      }
    
      input[type=submit]{
      width: 70%;
      background-color: rgb(178, 176, 187);
        border: none;
        color: rgb(1, 0, 11);
        padding: 16px 20px;
        text-decoration: none;
        margin: 4px 4px;
        cursor: pointer;
        border-radius: 25px;
      }
    
      h2{
        color: rgb(214, 219, 220);
        font-size: 35px;
        font-weight: 400;
        background-image: linear-gradient(to left, #fdfdff, #d1bce5);
      }
    
      .btn {

      background-color: white; 
      border: none; 
      color: white; 
      padding: 12px 16px; 
      font-size: 16px; 
      position: absolute;
    }
    
    button, h3{
      display: inline-block;
    }
    
    .btn:hover {
      background-color: rgb(255, 255, 255);
    }
    
    h3{
      font-weight: bold;
      color: rgb(0, 0, 128);
      padding-right: 33px;
      font-style: italic;
      margin-left: 9px;
      
    }
    h2 {
		color: #000000;
		text-transform: uppercase;
		font-size: 3em;
		font-weight: 100;
	}
    </style>