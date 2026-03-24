import * as React from 'react';

interface PaginationProps {
  totalPages: number;
  offset: number;
  onPageChange: (page: number) => void;
  siblingCount?: number;
}

const DOTS = '...';

const Pagination: React.FC<PaginationProps> = ({
  totalPages,
  offset,
  onPageChange,
  siblingCount = 1,
}) => {
  const currentPage = offset + 1;

  const getPageNumbers = (): (number | string)[] => {
    const totalPageNumbers = siblingCount * 2 + 5;

    if (totalPages <= totalPageNumbers) {
      return Array.from({ length: totalPages }, (_, i) => i + 1);
    }

    const leftSiblingIndex = Math.max(currentPage - siblingCount, 1);
    const rightSiblingIndex = Math.min(currentPage + siblingCount, totalPages);

    const showLeftDots = leftSiblingIndex > 2;
    const showRightDots = rightSiblingIndex < totalPages - 2;

    if (!showLeftDots && showRightDots) {
      const leftRange = Array.from({ length: siblingCount * 2 + 3 }, (_, i) => i + 1);
      return [...leftRange, DOTS, totalPages];
    }

    if (showLeftDots && !showRightDots) {
      const rightRange = Array.from({ length: siblingCount * 2 + 3 }, (_, i) => totalPages - (siblingCount * 2 + 2) + i);
      return [1, DOTS, ...rightRange];
    }

    if (showLeftDots && showRightDots) {
      const middleRange = Array.from({ length: siblingCount * 2 + 1 }, (_, i) => leftSiblingIndex + i);
      return [1, DOTS, ...middleRange, DOTS, totalPages];
    }

    return [];
  };

  const pageNumbers = getPageNumbers();

  const baseButton = 'flex h-9 w-9 items-center justify-center border border-black/10 bg-[color:var(--bg-surface)] font-primary text-[10px] uppercase tracking-[0.16em] text-[color:var(--text-muted)] transition-colors disabled:cursor-not-allowed disabled:opacity-40';
  const activeButton = 'border-black/12 bg-[color:var(--fg-strong)] text-white';
  const inactiveButton = 'hover:bg-[color:var(--bg-muted)] hover:text-[color:var(--text-primary)]';

  return (
    <div className="flex items-center gap-2">
      <button onClick={() => onPageChange(currentPage - 1)} disabled={currentPage === 1} className={`${baseButton} ${currentPage === 1 ? '' : inactiveButton}`}>
        &lt;
      </button>

      {pageNumbers.map((page, index) =>
        page === DOTS ? (
          <span key={index} className="px-2 font-primary text-[10px] uppercase tracking-[0.18em] text-[color:var(--text-muted)]">
            ...
          </span>
        ) : (
          <button key={index} onClick={() => onPageChange(Number(page))} className={`${baseButton} ${page === currentPage ? activeButton : inactiveButton}`}>
            {page}
          </button>
        )
      )}

      <button onClick={() => onPageChange(currentPage + 1)} disabled={currentPage === totalPages} className={`${baseButton} ${currentPage === totalPages ? '' : inactiveButton}`}>
        &gt;
      </button>
    </div>
  );
};

export default Pagination;
