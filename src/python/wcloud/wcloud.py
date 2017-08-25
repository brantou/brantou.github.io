import jieba.analyse
from wordcloud import WordCloud, ImageColorGenerator
import matplotlib.pyplot as plt
import numpy as np
from PIL import Image

font_path = '../resource/tyzkaishu.ttf'
width = 640
height = 480

text = open('../resource/xiyouji.txt').read()
words = jieba.analyse.extract_tags(text, topK=200, withWeight=True)

word_freqs = {}
for word in words:
    word_freqs[word[0]] = word[1]

wordcloud = WordCloud(
    font_path=font_path, width=width,
    height=height).generate_from_frequencies(word_freqs)
wordcloud.to_file('../images/xiyouji.png')

mask = np.array(Image.open('../resource/stormtrooper_mask.png'))
wordcloud = WordCloud(
    background_color="white",
    font_path=font_path,
    width=width,
    height=height,
    mask=mask).generate_from_frequencies(word_freqs)
wordcloud.to_file('../images/xiyouji-mask.png')

alice_coloring = np.array(Image.open('../resource/alice_color.png'))
# create coloring from image
image_colors = ImageColorGenerator(alice_coloring)
wc = WordCloud(
    font_path=font_path,
    width=width,
    height=height,
    background_color="white",
    mask=alice_coloring,
    max_font_size=40,
    random_state=42).generate_from_frequencies(word_freqs)

# recolor wordcloud and show
# we could also give color_func=image_colors directly in the constructor
plt.imshow(wc.recolor(color_func=image_colors), interpolation="bilinear")
plt.axis("off")
plt.figure()
plt.savefig('../images/xiyouji-color.png')
