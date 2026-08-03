describe(
  'Multi line describe',
  () => {
    it(
      `should support template literal names`,
      () => {
        // Given a template-literal test name
        // When the file is parsed
        // Then the scenario is produced
      }
    );

    // A stray comment that must NOT become a Gherkin step:
    //   this is just a note, not a Given/When/Then line.

    it('should ignore fit/xit siblings', () => {
      // Given the sibling calls are fit and xit
      // When the parser encounters them
      // Then they are not treated as scenarios
    });
  }
);

fit('should not become a scenario', () => {
  // Given fit is used
  // When the file is parsed
  // Then this block is ignored
});

xit('should also be ignored', () => {
  // Given xit is used
});
