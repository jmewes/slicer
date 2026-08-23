
Create a plan for the introduction of a new parser for "vitest"!

It should be structurally equivalent to the existing "javascript" parser (specialised on the Jasmine framework), but strictly separated.

## vitest grammar

(see https://vitest.dev/guide/learn/writing-tests.html and https://vitest.dev/guide/learn/writing-tests-with-ai.html for details)

For each of the vitest testsuite features listed below, create a separate test file and a separate test case.

### Simple tests

```ts
import { expect, test } from 'vitest'

test('foo bar baz', () => {

})
```

### Alias 'it' for 'test'

```ts
import { expect, it } from 'vitest'

it('foo bar baz', () => {

})
```

### Tests grouped with `describe`

```ts
import { describe, expect, test } from 'vitest'

describe('Spam eggs', () => {
  test('foo bar baz', () => {

  })
})
```

### afterEach / beforeEach


```ts
import { afterEach, beforeEach, expect, test } from 'vitest'

beforeEach(() => {

})

afterEach(() => {

})

test('foo bar baz', () => {

})

```

### Only one test

```ts
test.only('foo bar baz', () => {

})
```

### Skipping a test

```ts
test.skip('bar', () => {

})
```

### Test placeholder

```ts
test.todo('baz')
```

### Parameterised test

```ts
import { expect, test } from 'vitest'

test.for([
  [1, 2, 3],
])('foo bar baz', ([f, b, expected]) => {

})
```

### Async/Await 

```ts
import { expect, test } from 'vitest'

test('foo bar baz', async () => {

})
```

### Timeouts

```ts
test('foo bar baz', async () => {

}, 1_000)
```

### `afterAll` / `beforeAll`

```ts
import { afterAll, beforeAll, expect, test } from 'vitest'


beforeAll(async () => {

})

afterAll(async () => {

})

test('foo bar baz', async () => {

})
```

### Scoped setup

```ts
import { beforeEach, describe, expect, test } from 'vitest'

describe('spam eggs', () => {

  beforeEach(() => {

  })

  test('foo bar baz', () => {

  })

})
```

### Nested describe blocks

```ts
import { describe, expect, test } from "vitest";

describe("app-sections", () => {

  describe("getSectionMenuItems", () => {
    test("foo bar baz", () => {

    });
  });
});

```

## Test file naming conventions

The following test file pattern should be supported:

1. *.test.ts / *.test.js: e.g. utils.test.ts, utils.test.js
2. *.spec.ts / *.spec.js: Scaffold.spec.tsx, Scaffold.spec.ts

## Command-line flags

For the `slicer rev` command, add a flag `--parser` (short: `-p`) where one of the following options has to be provided `jasmine`, `vitest`.

## Documentation

### Slicer CLI usage

Add a basic user guide to the "usage" section of the README.

### Adding a new parser

In the file @docs/create-new-parser.md, add documentation for the steps that will need to be executed with later adding more parsers (e.g. Deno, JUnit).
