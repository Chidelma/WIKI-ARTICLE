<script lang="ts">
	import { onBeforeMount, onMounted, ref } from 'vue';

	export default  {

        props: {
			name: String
		},

		setup(props: any) {
            
            const article = ref('');

            onBeforeMount(async () => {

				const res: Response = await fetch('http://localhost:8080/articles/' + props.name, {
					method: 'GET',
					headers: {
						Accept: 'application/json',
					}
				});

				article.value = await res.text();
			});
            
            async function saveArticle() {

                const res: Response = await fetch('http://localhost:8080/articles/' + props.name, {
					method: 'PUT',
					headers: {
						Accept: 'application/json',
					},
					body: article.value
				});
            }

			return {
                article,
                saveArticle
			};
		}
	}
</script>

<template>
	<div>
		<h3 id="article">{{ name.toUpperCase() }}</h3>

		<textarea id="input-desc" class="desc" v-model="article" rows="4"></textarea>
		<p id="preview" class="desc">{{ article }}</p>

		<button class="btn"><router-link :to="{ name: 'Article', params: { name: name } }">Cancel</router-link></button>
		<button class="btn" @click="saveArticle()">Save</button>
	</div>
</template>

<style>
	.btn {
		margin-right: 20px;
	}

	textarea {
		width: 400px;
		height: 100px;
	}
</style>