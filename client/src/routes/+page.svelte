<script lang="ts">
	import { API_URL } from '#lib';

	import axios from 'axios';
	let url = $state('');
	let mode = $state<'text' | 'html' | 'json'>('text');

	const onScrape = async (event: SubmitEvent) => {
		event.preventDefault();

		// validate the form
		if (!url || !mode) console.error('Form validation failed');

		// validate the url inputted if the url format exist
		if (!url.startsWith('http://') && !url.startsWith('https://')) {
			console.error('Invalid URL format');
			return;
		}

		// try catch on posting the data
		try {
			const { data } = await axios.post(`${API_URL}/scrape`, { url, mode });
			console.log(data);
		} catch (error) {
			if (error && axios.isAxiosError(error)) {
				console.error(error.response?.data);
			}
		}

		//
	};
</script>

<svelte:head>
	<title>ScrapeMesh — Web Scraper</title>
	<meta name="description" content="Extract clean content from any webpage." />
</svelte:head>

<main class="flex min-h-screen items-center justify-center bg-[#09090b] px-4 text-zinc-100">
	<div class="w-full max-w-xl">
		<div class="mb-8 text-center">
			<div
				class="mx-auto mb-5 flex size-12 items-center justify-center rounded-xl border border-zinc-700 bg-zinc-900 text-xl font-bold text-white"
			>
				S
			</div>

			<h1 class="text-3xl font-semibold tracking-tight">
				ScrapeMesh<span class="text-blue-500">.</span>
			</h1>

			<p class="mt-3 text-sm text-zinc-400">Extract web content. Keep it simple.</p>
		</div>

		<form
			onsubmit={onScrape}
			class="rounded-2xl border border-zinc-800 bg-zinc-900/70 p-5 shadow-2xl sm:p-7"
		>
			<label for="url" class="mb-2 block text-sm font-medium text-zinc-300"> Target URL </label>

			<input
				id="url"
				type="url"
				bind:value={url}
				placeholder="https://example.com"
				required
				class="w-full rounded-lg border border-zinc-700 bg-zinc-950 px-4 py-3 text-sm transition outline-none placeholder:text-zinc-600 focus:border-blue-500 focus:ring-2 focus:ring-blue-500/20"
			/>

			<div class="mt-5">
				<label for="mode" class="mb-2 block text-sm font-medium text-zinc-300">
					Output format
				</label>

				<select
					id="mode"
					bind:value={mode}
					class="w-full rounded-lg border border-zinc-700 bg-zinc-950 px-4 py-3 text-sm outline-none focus:border-blue-500"
				>
					<option value="text">Plain text</option>
					<option value="html">HTML</option>
					<option value="json">JSON</option>
				</select>
			</div>

			<button
				type="submit"
				class="mt-6 w-full rounded-lg bg-blue-600 px-4 py-3 text-sm font-medium text-white transition hover:bg-blue-500 active:scale-[0.99]"
			>
				Start scraping →
			</button>

			<p class="mt-4 text-center text-xs text-zinc-500">
				Choose a URL and output format to get started.
			</p>
		</form>

		<p class="mt-6 text-center text-xs text-zinc-600">ScrapeMesh · Web content extraction</p>
	</div>
</main>
