import type { ReactNode } from 'react';
import { DocsLayout } from 'fumadocs-ui/layouts/docs';
import { Github } from 'lucide-react';
import { source } from '@/lib/source';

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <DocsLayout
      tree={source.pageTree}
      nav={{
        title: 'sarvamai-go SDK Documentation',
        url: '/docs',
        transparentMode: 'top',
        children: (
          <div data-top-actions>
            <a
              href="https://github.com/Shreehari-Acharya/sarvamai-go"
              target="_blank"
              rel="noreferrer noopener"
              aria-label="GitHub Repository"
              title="GitHub Repository"
              data-top-github
            >
              <Github size={16} strokeWidth={2} aria-hidden="true" />
            </a>
          </div>
        ),
      }}
      sidebar={{
        defaultOpenLevel: 1,
        collapsible: true,
      }}
      themeSwitch={{
        enabled: false,
      }}
    >
      {children}
    </DocsLayout>
  );
}
