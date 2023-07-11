
- ### jieba分词结果输入BERT

 ```
 def bert_chinese(words):
    tokenizer = BertTokenizer.from_pretrained('bert-base-chinese')
    model = BertModel.from_pretrained('bert-base-chinese')
    output_list = []
    for text in words:
        # 将文本转化为Bert可以接受的格式
        input_ids = torch.tensor([tokenizer.encode(text, add_special_tokens=True)])
        # 获取BERT模型的输出
        outputs = model(input_ids)
        # 取第1层的输出，即对应的词向量
        output = outputs[0][0][1:-1, :]
        # 添加到输出列表
        output_list.append(output)
        # 它是对一个三维数组（outputs）的第一个元素（outputs[0]）进行切片，取出第一维中除去第一个和最后一个元素以外的所有元素（outputs[0][0][1:-1, :]）。
        # 其中，1:-1表示从第二个到倒数第二个元素（不包括第一个和最后一个），
        # 而":"则表示取出所有的第二维元素。因此，这一操作是将原数组中第一个维度上除去两端元素以外的所有部分保留下来，并返回一个新的numpy数组。
    # print(output_wordVec[0])
    # 字向量
    # print(output_wordVec[1])
    # 句向量
    # 返回值是个元祖，第一个内容是last_hidden_state=tensor...，第二个是 pooler_output=tensor...,
    # 然后是hidden_states=None, past_key_values=None, attentions=None, cross_attentions=None
    # Batch size outputs = model(input_ids)
    # sequence_output = outputs[0]
    # pooled_output = outputs[1]
    # print(sequence_output.shape)
    # # torch.Size([1, 6, 768]) 我是学生是4个字为什么会有6个字的向量【cls】【sqp】
    # # 字向量sequence_output[0][1]就是“我”的字向量
    # print(pooled_output.shape)

```

- 加载BERT模型和tokenizer
- 输入列表list
- text_list = ['今天天气不错', '我喜欢吃苹果', '北京是中国的首都']
- 分别对每个词语进行BERT编码并获取输出向量
- 在上面的代码中，首先加载预训练的BERT模型和tokenizer。
- 然后遍历输入的词语列表，对每个词语进行编码并获取BERT模型的输出。
- 这里使用的是BERT的第1层输出，即对应的词向量。
- 最后将每个词的输出向量添加到输出列表中，打印输出列表即可。

### docx2pdf
```
import docx2pdf
import os
```

- 文件名不要叫包名！！！
- AttributeError: partially initialized module 'docx2pdf' has no attribute 'convert'

```
# document = Document('example.docx')
# document.save('example.pdf')
```

- 这将使用docx模块打开名为example.docx的文件，并使用document.save()命令将其另存为example.pdf文件。
- 但是，此方法将生成一个新的pdf文件，而不是将docx文件转换为现有的pdf文件。
- 如果你想将docx文件转换为pdf，而不是生成另一个文件的话，需要使用docx2pdf模块。以下是相应代码：
```
filename = 'data/5.docx'
file_basename = os.path.basename(filename)
```
- 在Python中，可以使用字符串操作来分离文件名的前缀和后缀。可以使用 `os.path.splitext()`函数将文件名和扩展名分开。
- `os.path.splitext()`函数会将文件名和扩展名以元组的形式返回。
`
file_prefix, file_suffix = os.path.splitext(file_basename)`

- 在上面的示例代码中，`os.path.splitext()`函数返回了`('example', '.txt')`元组
```
docx2pdf.convert("data/5.docx", "data/docx2pdf/"+file_prefix+".pdf")
```

- 这将使用docx2pdf模块将example.docx转换为example.pdf。运行此代码段时，需要安装docx2pdf模块。您可以使用pip安装docx2pdf：
- 在上面的代码中，`os.path.basename()`函数可以获取指定路径的文件名（或文件夹名）。
- 如果您只需要获取文件名而不需要路径，那么您可以先用`os.path.dirname()`函数获取文件所在的文件夹路径，再使用`os.path.basename()`函数获取文件名。

