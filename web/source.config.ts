import { defineDocs, frontmatterSchema } from 'fumadocs-mdx/config';

export const docs = defineDocs({
  dir: '../docs',
  docs: {
    schema: frontmatterSchema,
  },
});
