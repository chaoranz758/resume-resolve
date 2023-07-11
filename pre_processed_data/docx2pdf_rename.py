import docx2pdf
import os


def trans_docx(a):
    file_basename = os.path.basename(a)
    file_prefix, file_suffix = os.path.splitext(file_basename)
    if file_suffix == '.docx':
        docx2pdf.convert(a, "data/docx2pdf/" + file_prefix + ".pdf")


def main_trans():
    filepath = 'data/13.docx'
    trans_docx(filepath)


main_trans()