let userHasInteracted = false;

if (typeof window !== 'undefined') {
  window.addEventListener(
    'click',
    () => {
      userHasInteracted = true;
    },
    { once: true }
  );
}

const notify = (audioUrl) => {
  if (typeof audioUrl !== 'string') {
    console.error("Invalid audio path provided");
    return;
  }

  const audio = new Audio(audioUrl);

  if (userHasInteracted) {
    audio.play().catch((error) => {
      console.error("Audio playback failed:", error);
    });
  } else {
    console.warn("Audio not played: user hasn't interacted with the page yet");
  }

  // Reset audio after it ends
  audio.addEventListener('ended', () => {
    audio.currentTime = 0;
  });
};

export default notify;
