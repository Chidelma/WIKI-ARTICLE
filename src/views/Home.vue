<script lang="ts">
	import { onBeforeMount, ref } from 'vue';

	export default  {
		setup() {

			const articles = ref([]);

			onBeforeMount(async () => {

				const res: Response = await fetch('http://localhost:8080/articles/', {
					method: 'GET',
					headers: {
						Accept: 'application/json',
					}
				});

				articles.value = await res.json();
			});

			return {
				articles,
			};
		}
	}
</script>

<template>
	<div>
		<h3>ARTICLES</h3>
		<div class="article" v-for="article in articles" :key="article">
			<button class="btn"><router-link :to="{ name: 'Article', params: { name: article } }">{{ article }}</router-link></button>
		</div>
	</div>
</template>

<style>
    .btn {
        margin-right:20px;
    }
</style>