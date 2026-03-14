/**
 * Svelte action that animates text with a "decrypt/descramble" typewriter effect.
 *
 * Text is revealed left-to-right. At the leading edge, a few characters cycle
 * through random glyphs before "locking in" to the real character. Spaces are
 * preserved (never scrambled) so word shapes stay readable.
 *
 * Usage:
 *   <!-- Always animate -->
 *   <span use:decryptText={"hello"}></span>
 *
 *   <!-- Skip animation on mount, animate subsequent updates -->
 *   <span use:decryptText={{ text: "hello", immediate: true }}></span>
 *
 * The element's textContent is fully controlled by the action — leave the
 * element body empty in the template.
 */

const GLYPHS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%&*<>{}[]';

const CHAR_INTERVAL_MS = 20;
const SCRAMBLE_WIDTH = 4;

export type DecryptTextInput = string | { text: string; immediate?: boolean };

function unpack(input: DecryptTextInput): { text: string; immediate: boolean } {
	if (typeof input === 'string') return { text: input, immediate: false };
	return { text: input.text, immediate: input.immediate ?? false };
}

function randomGlyph(): string {
	return GLYPHS[Math.floor(Math.random() * GLYPHS.length)];
}

export function decryptText(node: HTMLElement, input: DecryptTextInput) {
	let frameId: number | null = null;
	let target = '';
	let revealed = 0;
	let lastRevealTime = 0;

	function render() {
		let display = target.slice(0, revealed);
		const end = Math.min(revealed + SCRAMBLE_WIDTH, target.length);
		for (let i = revealed; i < end; i++) {
			display += target[i] === ' ' ? ' ' : randomGlyph();
		}
		node.textContent = display;
	}

	function tick(time: number) {
		if (!lastRevealTime) lastRevealTime = time;

		const elapsed = time - lastRevealTime;
		if (elapsed >= CHAR_INTERVAL_MS) {
			const chars = Math.floor(elapsed / CHAR_INTERVAL_MS);
			revealed = Math.min(revealed + chars, target.length);
			lastRevealTime += chars * CHAR_INTERVAL_MS;
		}

		render();

		if (revealed < target.length) {
			frameId = requestAnimationFrame(tick);
		}
	}

	function start(newTarget: string, fromIndex: number) {
		if (frameId !== null) cancelAnimationFrame(frameId);
		target = newTarget;
		revealed = fromIndex;
		lastRevealTime = 0;
		if (fromIndex < newTarget.length) {
			frameId = requestAnimationFrame(tick);
		} else {
			node.textContent = newTarget;
		}
	}

	const { text, immediate } = unpack(input);
	if (immediate) {
		target = text;
		revealed = text.length;
		node.textContent = text;
	} else {
		start(text, 0);
	}

	return {
		update(newInput: DecryptTextInput) {
			const { text: newText } = unpack(newInput);
			if (newText === target && revealed >= target.length) return;

			// Find common prefix so we don't re-animate already-revealed text
			const stableLen = Math.min(revealed, target.length);
			let common = 0;
			while (common < stableLen && common < newText.length && target[common] === newText[common]) {
				common++;
			}

			start(newText, common);
		},
		destroy() {
			if (frameId !== null) cancelAnimationFrame(frameId);
		}
	};
}
