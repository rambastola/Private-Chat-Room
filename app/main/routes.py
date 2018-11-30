from flask import session, redirect, url_for, render_template, request, abort, Response
from . import main
from .forms import LoginForm
from .forms import MainLogin
import requests

@main.route('/', methods=['GET', 'POST'])
def authenticate():
    """make sure users' log in"""
    form = MainLogin()

    if form.validate_on_submit():
        session['username'] = form.username.data
        session['password'] = form.password.data
        if session['username'] == "ram" and session['password'] == "berea":
            return redirect(url_for('.index'))
        elif session['username'] == "jamal" and session['password'] == "berea":
            return redirect(url_for('.index'))
        elif session['username'] == "basant" and session['password'] == "berea":
            return redirect(url_for('.index'))

    elif request.method == 'GET':
        form.username.data = session.get('username', '')
        form.password.data = session.get('password', '')

    return render_template('login.html', form=form)

@main.route('/chatsession', methods=['GET', 'POST'])
def index():
    """Login form to enter a room."""
    form = LoginForm()
    if form.validate_on_submit():
        session['name'] = form.name.data
        session['room'] = form.room.data
        return redirect(url_for('.chat'))
    elif request.method == 'GET':
        form.name.data = session.get('name', '')
        form.room.data = session.get('room', '')
    return render_template('index.html', form=form)

@main.route('/chat')
def chat():
    """Chat room. The user's name and room must be stored in
    the session."""
    name = session.get('name', '')
    room = session.get('room', '')
    if name == '' or room == '':
        return redirect(url_for('.index'))
    return render_template('chat.html', name=name, room=room)
