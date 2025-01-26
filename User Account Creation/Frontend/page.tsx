'use client';

import { useState } from 'react';
import { Eye, EyeOff } from 'lucide-react';
import { Button } from '../components/ui/button';
import { Input } from '../components/ui/input';
import { Checkbox } from '../components/ui/checkbox';
import { Label } from '../components/ui/label';

export default function AuthPage() {
  const [isLogin, setIsLogin] = useState(true);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [name, setName] = useState('');
  const [showPassword, setShowPassword] = useState(false);

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center p-4">
      <div className="w-full max-w-md bg-white rounded-lg shadow-lg p-8">
      <div className="flex justify-center mb-6">
  <img
    src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTarH0zgXRoEHwRY1w53vhDE4uFxD-EeqfdhEQ_DdJGnR28euAKbC62HLozd8-ISEIDow&usqp=CAU"
    alt="Financfy Logo"
    className="w-20 h-30"
  />
</div>

        {/* Toggle Tabs */}
        <div className="flex justify-center mb-6">
          <button
            onClick={() => setIsLogin(true)}
            className={`px-4 py-2 text-sm font-medium ${
              isLogin ? 'text-teal-700 border-b-2 border-teal-700' : 'text-gray-500'
            }`}
          >
            Login
          </button>
          <button
            onClick={() => setIsLogin(false)}
            className={`px-4 py-2 text-sm font-medium ${
              !isLogin ? 'text-teal-700 border-b-2 border-teal-700' : 'text-gray-500'
            }`}
          >
            Register
          </button>
        </div>

        {/* Forms */}
        {isLogin ? (
          <form className="space-y-6">
            <div>
              <Label htmlFor="email" className="block text-sm text-teal-700">
                Email
              </Label>
              <Input
                id="email"
                type="email"
                placeholder="Enter your email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full mt-1"
                required
              />
            </div>
            <div>
              <Label htmlFor="password" className="block text-sm text-teal-700">
                Password
              </Label>
              <div className="relative">
                <Input
                  id="password"
                  type={showPassword ? 'text' : 'password'}
                  placeholder="Enter your password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full mt-1"
                  required
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute right-3 top-1/2 transform -translate-y-1/2 text-teal-700 hover:text-teal-900"
                >
                  {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
                </button>
              </div>
            </div>
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <Checkbox id="remember" />
                <Label htmlFor="remember">Remember me</Label>
              </div>
              <a href="#" className="text-sm text-teal-700 hover:underline">
                Forgot password?
              </a>
            </div>
            <Button type="submit" className="w-full py-2 bg-teal-700 text-white rounded-md hover:bg-teal-800">
              Login
            </Button>
          </form>
        ) : (
          <form className="space-y-6">
            <h1 className="text-2xl font-semibold text-center text-teal-700">Create an Account</h1>
            <div>
              <Label htmlFor="name" className="block text-sm text-teal-700">
                Full Name
              </Label>
              <Input
                id="name"
                type="text"
                placeholder="Enter your full name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="w-full mt-1"
                required
              />
            </div>
            <div>
              <Label htmlFor="email" className="block text-sm text-teal-700">
                Email
              </Label>
              <Input
                id="email"
                type="email"
                placeholder="Enter your email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full mt-1"
                required
              />
            </div>
            <div>
              <Label htmlFor="password" className="block text-sm text-teal-700">
                Password
              </Label>
              <div className="relative">
                <Input
                  id="password"
                  type={showPassword ? 'text' : 'password'}
                  placeholder="Enter your password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full mt-1"
                  required
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute right-3 top-1/2 transform -translate-y-1/2 text-teal-700 hover:text-teal-900"
                >
                  {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
                </button>
              </div>
            </div>
            <div>
              <Label htmlFor="confirm-password" className="block text-sm text-teal-700">
                Confirm Password
              </Label>
              <Input
                id="confirm-password"
                type={showPassword ? 'text' : 'password'}
                placeholder="Confirm your password"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                className="w-full mt-1"
                required
              />
            </div>
            <div className="flex items-center">
              <Checkbox id="terms" />
              <Label htmlFor="terms" className="ml-2 text-sm text-teal-700">
                I agree to the{' '}
                <a href="#" className="font-medium text-teal-900 hover:underline">
                  terms and conditions
                </a>
              </Label>
            </div>
            <Button type="submit" className="w-full py-2 bg-teal-700 text-white rounded-md hover:bg-teal-800">
              Register
            </Button>
          </form>
        )}
      </div>
    </div>
  );
}
