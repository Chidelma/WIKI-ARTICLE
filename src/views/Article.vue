<script lang="ts">
	import { onBeforeMount, ref } from 'vue';
	import NotFound from './NotFound.vue';

	export default  {

        props: {
			name: String
		},

		setup(props: any) {

			const desc = ref('');

			onBeforeMount(async () => {

				const res: Response = await fetch('http://localhost:8080/articles/' + props.name, {
					method: 'GET',
					headers: {
						Accept: 'application/json',
					}
				});

				desc.value = await res.text();
			});

			return {
				desc,
			};
		}
	}
</script>

<template>
	<div>
		<button class="btn"><router-link to="/">Home</router-link></button>
		<button class="btn"><router-link :to="{ name: 'Edit', params: { name: name } }">Edit</router-link></button>

		<h3 id="article" v-if="desc != ''">{{ name.toUpperCase() }}</h3>

		<div v-if="desc != ''">
			<p id="article-desc">{{ desc }}</p>
		</div>
		<div v-else>
			<h5 id="not-found">No article with this exact name found. Use Edit button in the header to add it</h5>
		</div>
	</div>
</template>

<style>
	.btn {
		margin-right:20px;
	}
</style>