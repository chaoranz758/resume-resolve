
import json
import jieba
import jieba.analyse
import PyPDF2
from hanlp import *
import re

def extract_pdf():
    with open(r'./data./docx2pdf/5.pdf', 'rb') as f:
        pdf_reader = PyPDF2.PdfReader(f)
        # 没有解决换行问题!
        text = ''
        for page in pdf_reader.pages:
            text += page.extract_text()
        retext = text
        text = re.sub(r'\n', '', text)

        return text,retext


stopwords = ['的', '是', '在', '这', '这个', ' ', '#']  # 自定义停用词表


def remove_stopwords(words):
    filtered_words = []
    for word in words:
        if word not in stopwords:
            filtered_words.append(word)
    return filtered_words


def tokenize(text):
    jieba.load_userdict('userdict.txt')
    words = jieba.cut(text, HMM=True)
    words = list(words)
    return words


# def bert_chinese(words):
#     tokenizer = BertTokenizer.from_pretrained('bert-base-chinese')
#     model = BertModel.from_pretrained('bert-base-chinese')
#     output_list = []
#     for text in words:
#         # 将文本转化为Bert可以接受的格式
#         input_ids = torch.tensor([tokenizer.encode(text, add_special_tokens=True)])
#         # 获取BERT模型的输出
#         outputs = model(input_ids)
#         # 取第1层的输出，即对应的词向量
#         output = outputs[0][0][1:-1, :]
#         # 添加到输出列表
#         output_list.append(output)
#         return output_list


def init_main():
    text,retext = extract_pdf()
    words = tokenize(text)
    filtered_words = remove_stopwords(words)
    # output_BERT = bert_chinese(filtered_words)
    # print(output_BERT[0].shape)
    return text,retext