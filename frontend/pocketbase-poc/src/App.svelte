<script>

import PocketBase from 'pocketbase';
const client = new PocketBase('http://127.0.0.1:8090');

let resultData = [];

async function loginUser() {
	const authData = await client.users.authViaEmail("user1@test.com", "Adisunny123");

	if (authData) {
		getAllRecords();
	}

	console.log("resultData: ", resultData)
}

loginUser();

async function createRecord() {
	console.log("create a record");
	const data = {
		name: "Test user",
	}
	const record = await client.records.create('demo', data);
}

async function getAllRecords() {
	console.log("reading all data");
	const resultList = await client.records.getList('demo', 1, 50, {
	    filter: 'created >= "2022-01-01 00:00:00"',
	});
	console.log(resultList);

	resultData = resultList.items;
}

async function updateRecord(id) {
  console.log("upating record");

  const data = {
    name: "Test user 2",
  };
  const record = await client.records.update('demo', id, data);
}

async function deleteRecord(id) {
  console.log("delete record");

  const record = await client.records.delete('demo', id);
}


</script>

<div>
  Hello Pocketbase
</div>

{#each resultData as r} 
  <h1>{r.name}</h1>
{/each}

<button on:click={createRecord}> 
  Add
</button>

<button on:click={() => {updateRecord("xeflgvrzocpotbb")}}>
  Update
</button>

<button on:click={() => {deleteRecord("xeflgvrzocpotbb")}}>
  Delete
</button>