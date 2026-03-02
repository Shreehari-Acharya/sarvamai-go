import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { DocsBody, DocsDescription, DocsPage, DocsTitle, EditOnGitHub } from 'fumadocs-ui/page';
import defaultMdxComponents from 'fumadocs-ui/mdx';
import { source } from '@/lib/source';

const GITHUB_OWNER = 'Shreehari-Acharya';
const GITHUB_REPO = 'sarvamai-go';
const GITHUB_REF = 'main';

type PageProps = {
  params: Promise<{ slug?: string[] }>;
};

export default async function DocPage({ params }: PageProps) {
  const resolvedParams = await params;
  const page = source.getPage(resolvedParams.slug);

  if (!page) {
    notFound();
  }

  const MdxContent = page.data.body;
  const filePath =
    page.path ?? `${resolvedParams.slug?.length ? resolvedParams.slug.join('/') : 'index'}.mdx`;
  const editPath = filePath.startsWith('docs/') ? filePath : `docs/${filePath}`;
  const editUrl = `https://github.com/${GITHUB_OWNER}/${GITHUB_REPO}/blob/${GITHUB_REF}/${editPath}`;

  return (
    <DocsPage toc={page.data.toc} full={page.data.full}>
      <div className='w-10'>
        <EditOnGitHub href={editUrl} />
      </div>
      <DocsTitle>{page.data.title}</DocsTitle>
      <DocsDescription>{page.data.description}</DocsDescription>
      <DocsBody>
        <MdxContent components={defaultMdxComponents} />
      </DocsBody>
    </DocsPage>
  );
}

export async function generateStaticParams() {
  return source.generateParams();
}

export async function generateMetadata({ params }: PageProps): Promise<Metadata> {
  const resolvedParams = await params;
  const page = source.getPage(resolvedParams.slug);

  if (!page) {
    notFound();
  }

  return {
    title: page.data.title,
    description: page.data.description,
  };
}
