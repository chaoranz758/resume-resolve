import re

import torch
from transformers import PegasusTokenizer, PegasusForConditionalGeneration

# 加载预训练的PEGASUS模型和分词器
model_name = 'google/pegasus-large'  # 模型名称
tokenizer = PegasusTokenizer.from_pretrained(model_name)
model = PegasusForConditionalGeneration.from_pretrained(model_name)

# 待摘要的文本
text = "自我学习能力还是比较强的，想做的事很认真。专业知识扎实，有积极的工作态度，能够独立工作，又有团队精神。具有良好的文化素质，在未来的工作中，我将以充沛的精力，努力工作，稳定地进步自己的工作能力。我正在寻找一个更好的发展平台，希望能够充分发挥自己的优势，共同努力成就一番事业 "
text = re.sub(r'，', '', text)
text = re.sub(r'。', '', text)
# 对文本进行编码
inputs = tokenizer.encode(text, return_tensors='pt')

# 生成摘要
summary_ids = model.generate(inputs, max_length=50, num_beams=4, early_stopping=True)
summary = tokenizer.decode(summary_ids[0], skip_special_tokens=True)

# 打印摘要结果
print("摘要：", summary)


