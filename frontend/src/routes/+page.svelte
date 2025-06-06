<script lang="ts">
	import { apiCall } from '$lib/utils/api';
	import { validateField } from '$lib/utils/validate';
	import { passwordSchema, usernameSchema, emailSchema } from '$lib/utils/validators';
	import { triggerError, triggerInfo } from '../stores/modal';
	import { Eye, EyeOff } from 'lucide-svelte';
	let showErrorModal: boolean = false;
	let isSubmitting: boolean = false;
	let showPassword: boolean = false;
	let errorMessage: string | null = '';

	let formData: { username: string; password: string } = {
		username: '',
		password: ''
	};

	const handleSubmit = async (): Promise<void> => {
		errorMessage = '';
		const usernameError = validateField(usernameSchema, formData.username);
		const emailError = validateField(emailSchema, formData.username);
		if (!(!usernameError || !emailError)) {
			errorMessage = usernameError || emailError;
			return;
		}

		const passwordError = validateField(passwordSchema, formData.password);
		if (passwordError) {
			errorMessage = passwordError;
			return;
		}

		isSubmitting = true;
		try {
			await apiCall<any>({
				endpoint: '/v1/auth/login',
				method: 'POST',
				data: formData,
				credentials: 'include',
				onSuccess: (response) => {
					localStorage.setItem('user', JSON.stringify(response.user));
					triggerInfo('Login Successful');
				}
			});
		} catch (error) {
			triggerError('An error occurred, please try again later.');
		} finally {
			isSubmitting = false;
		}
	};
	const closeModal = () => {
		showErrorModal = false;
	};
</script>

<main
	class="flex min-h-screen items-center justify-center bg-gradient-to-br from-gray-800 via-gray-900 to-black px-4"
>
	<div
		class="w-full max-w-md rounded-3xl bg-gray-800/90 p-10 shadow-xl backdrop-blur-md transition-all duration-300 hover:shadow-2xl"
	>
		<div class="mb-6 text-center">
			<h1 class="text-4xl font-bold text-indigo-500">BuzzChat</h1>
			<p class="mt-1 text-sm text-gray-400">Welcome back! Log in to start messaging.</p>
		</div>

		<form on:submit|preventDefault={handleSubmit} class="space-y-4">
			<div>
				<label for="username" class="block pb-0.5 pl-1 text-sm font-medium text-gray-300"
					>Username/Email</label
				>
				<input
					type="text"
					id="username"
					bind:value={formData.username}
					placeholder="Username/Email"
					class="w-full rounded-xl border border-gray-600 bg-gray-700 px-4 py-2 text-white shadow-sm focus:border-indigo-500 focus:ring-2 focus:ring-indigo-400 focus:outline-none"
					required
				/>
			</div>

			<div>
				<label for="password" class="block pb-0.5 pl-1 text-sm font-medium text-gray-300"
					>Password</label
				>
				<div class="relative">
					<input
						type={showPassword ? 'text' : 'password'}
						id="password"
						bind:value={formData.password}
						placeholder="Password"
						class="w-full rounded-xl border border-gray-600 bg-gray-700 px-4 py-2 pr-12 text-white shadow-sm focus:border-indigo-500 focus:ring-2 focus:ring-indigo-400 focus:outline-none"
						required
					/>
					<button
						type="button"
						on:click={() => (showPassword = !showPassword)}
						class="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-400 transition-colors duration-200 hover:cursor-pointer hover:text-indigo-400 focus:outline-none"
						aria-label={showPassword ? 'Hide password' : 'Show password'}
					>
						{#if showPassword}
							<EyeOff class="h-5 w-5" />
						{:else}
							<Eye class="h-5 w-5" />
						{/if}
					</button>
				</div>
			</div>

			<!-- Forgot Password Link Above Login Button -->
			<div class="text-right text-sm">
				<a href="/forgot-password" class="text-indigo-500 hover:underline">Forgot your password?</a>
			</div>

			{#if errorMessage !== ''}
				<span class="ml-2 block text-sm text-red-500">{errorMessage}</span>
			{/if}

			<button
				type="submit"
				disabled={isSubmitting}
				class="group relative flex w-full items-center justify-center gap-2 rounded-xl bg-indigo-600 px-4 py-2 font-medium text-white transition hover:cursor-pointer hover:bg-indigo-700 focus:ring-2 focus:ring-indigo-500 disabled:cursor-default disabled:bg-gray-600"
			>
				{#if isSubmitting}
					<span class="h-5 w-5 animate-spin rounded-full border-2 border-white border-t-transparent"
					></span>
					<span></span>
				{:else}
					<span>Login</span>
				{/if}
			</button>
		</form>

		<div class="mt-6 text-center text-sm text-gray-400">
			Don't have an account?
			<a href="/signup" class="text-indigo-500 hover:underline">Sign up</a>
		</div>
	</div>
</main>
