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

    const firstPage = 1;
    const lastPage = totalPages;

    if (!showLeftDots && showRightDots) {
      const leftRange = Array.from(
        { length: siblingCount * 2 + 3 },
        (_, i) => i + 1
      );
      return [...leftRange, DOTS, totalPages];
    }

    if (showLeftDots && !showRightDots) {
      const rightRange = Array.from(
        { length: siblingCount * 2 + 3 },
        (_, i) => totalPages - (siblingCount * 2 + 2) + i
      );
      return [firstPage, DOTS, ...rightRange];
    }

    if (showLeftDots && showRightDots) {
      const middleRange = Array.from(
        { length: siblingCount * 2 + 1 },
        (_, i) => leftSiblingIndex + i
      );
      return [firstPage, DOTS, ...middleRange, DOTS, lastPage];
    }

    return [];
  };

  const pageNumbers = getPageNumbers();

  const baseButton =
    'h-9 w-9 border border-border flex items-center justify-center text-[10px] font-bold uppercase tracking-[0.15em] text-secondary transition-colors disabled:text-muted disabled:cursor-not-allowed disabled:border-border/40 rounded-none';
  const activeButton = 'bg-fluxus-primary text-black border-fluxus-primary';
  const inactiveButton = 'hover:bg-bg-card-hover';

  return (
    <div className="flex items-center gap-2">
      <button
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage === 1}
        className={`${baseButton} ${currentPage === 1 ? '' : inactiveButton}`}
      >
        &lt;
      </button>

      {pageNumbers.map((page, index) =>
        page === DOTS ? (
          <span key={index} className="px-2 text-secondary text-xs tracking-[0.2em]">
            ...
          </span>
        ) : (
          <button
            key={index}
            onClick={() => onPageChange(Number(page))}
            className={`${baseButton} ${page === currentPage ? activeButton : inactiveButton}`}
          >
            {page}
          </button>
        )
      )}

      <button
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage === totalPages}
        className={`${baseButton} ${currentPage === totalPages ? '' : inactiveButton}`}
      >
        &gt;
      </button>
    </div>
  );
};

export default Pagination;
