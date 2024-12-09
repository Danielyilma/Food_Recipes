export const getPrepTimeMinutes = (prepTime: string): number => {
  const minutes = parseInt(prepTime);
  return isNaN(minutes) ? 0 : minutes;
};

export const formatPrepTime = (minutes: number): string => {
  return `${minutes} mins`;
};
