from turtle import back
import google.generativeai as genai
from colorama import init, Fore, Back, Style
import os


# Initialize colorama
init(autoreset=True)

genai.configure(api_key=os.environ["API_KEY"])

model = genai.GenerativeModel('gemini-1.5-flash')

def generate_prompt(library_name, vulnerability, brief_description):
    return f'''summarize the following text, It's important to patch these vulnerabilities to protect systems from attacks.
                "The format of the first sentence must be of {library_name} is vulnerable to {vulnerability}.
                 the second sentence must be of the format of "the vulnerability is due to {vulnerability}  due to {brief_description}"
              '''
# Example usage
library_name = "@strapi/plugin-upload"
vulnerability = "Denial-of-Service (DoS)"
brief_description = """A Denial-of-Service was found in the media upload process causing the server to crash without restarting,
affecting either development and production environments.DetailsUsually, 
errors in the application cause it to log the error and keep it running for other clients. 
This behavior, in contrast, stops the server execution, making it unavailable for any clients until it's manually restarted."""


prompt = generate_prompt(library_name,vulnerability,brief_description)
response = model.generate_content(prompt)

print(f"{Fore.CYAN}Prompt: {Fore.RED}{prompt}")

print(f"{Fore.CYAN}Generated description: {Fore.GREEN}{response.text}")


