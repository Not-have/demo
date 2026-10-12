import gradio as gr

def str(val):
  return val + '测试'

def reverse_text(text):
  return text[::-1]

demo = gr.Interface(
  fn=str,
  inputs="text",
  outputs="text"
)

demo.launch()
